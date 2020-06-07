package main

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// CreateWebsocket creates a new websocket connection to Discord
func (c *Client) CreateWebsocket() error {

	var err error

	// Create Websocket
	header := http.Header{}
	header.Add("accept-encoding", "zlib")
	c.wsConn, _, err = websocket.DefaultDialer.Dial(c.GatewayURL, header)
	if err != nil {
		return err
	}
	c.wsConn.SetCloseHandler(func(code int, text string) error {
		c.Connected = nil
		c.Log.WithField("Websocket closed:", code)
		return nil
	})

	AddEventHandlers()

	// Handle Hello
	err = c.Hello()
	if err != nil {
		return err
	}

	// Identify/Login
	err = c.Login()
	if err != nil {
		return err
	}

	// Start listening and heartbeat
	c.Connected = make(chan interface{})

	go c.Heartbeat(c.wsConn)
	go c.Listen(c.wsConn)

	c.Log.Info("Websocket: created")
	return nil
}

// Listen retrieves messages from the websocket
func (c *Client) Listen(wsConn *websocket.Conn) {
	c.Log.Info("Websocket: started listening")

	for {
		// Read message from connection
		messageType, message, err := wsConn.ReadMessage()
		if err != nil {
			c.Log.Errorf("Websocket: was unable to read message: %v", err)
		}

		// If closed close connection
		if messageType == -1 {
			c.wsConn.Close()
			c.Connected <- nil
			break
		}

		// Resolve event
		_, err = c.ResolveEvent(messageType, message)
		if err != nil {
			c.Log.Errorf("Websocket: Error resolving event: %v", err)
		}

		select {
		case <-c.Connected:
			c.Log.Info("Websocket: stopped listening")
			break
		default:
		}
	}
}

// ResolveEvent decodes a message and resolves the Event within it
func (c *Client) ResolveEvent(messageType int, message []byte) (*Event, error) {

	var reader io.Reader
	reader = bytes.NewBuffer(message)

	if messageType == websocket.BinaryMessage {
		z, _ := zlib.NewReader(reader)
		reader = z
	}

	var e *Event
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&e)
	if err != nil {
		c.Log.Error(messageType, message)
		return e, err
	}

	if e.Opcode == opcodeInvalidSession {
		reader = bytes.NewBuffer(message)

		if messageType == websocket.BinaryMessage {
			z, _ := zlib.NewReader(reader)
			reader = z
		}

		var eventInvalidSession *EventInvalidSession
		decoder := json.NewDecoder(reader)
		err := decoder.Decode(&eventInvalidSession)
		if err != nil {
			c.Log.Error(messageType, message)
			return e, err
		}
	}

	// Store sequence
	atomic.StoreInt64(c.Sequence, e.Sequence)

	// Do opcode specific things
	if e.Opcode == opcodeHeartbeatACK {
		c.LastHeartbeatAck = time.Now()
		c.Log.Debug("Websocket: received opcode 11 HeartbeatACK.")
		return e, nil
	}

	if e.Opcode == opcodeDispatch {
		c.Log.Debugf("Websocket: received opcode 0 Dispatch with event %s from Discord.", e.Type)
		e.Struct = eventHandlers[e.Type].New()

		if err = json.Unmarshal(e.RawData, &e.Struct); err != nil {
			c.Log.Errorf("error unmarshalling %s event, %s", e.Type, err)
		}
		if e.Type == messageCreateEventType {
			t := e.Struct.(*EventMessageCreate)
			c.Log.WithFields(logrus.Fields{
				"author":   t.Author.Username,
				"message":  t.Content,
				"serverid": t.GuildID,
			}).Debug("Websocket: received message")
		}
		return e, nil
	}

	c.Log.Debugf("Websocket: received opcode %v %v from Discord.", e.Opcode, e.Opcode.String())
	return e, nil

}

// Heartbeat sends heartbeat payloads to discord to signal discord that the client is still alive
func (c *Client) Heartbeat(wsConn *websocket.Conn) {

	// Set up ticker
	ticker := time.NewTicker(c.HeartbeatInterval)
	defer ticker.Stop()

	for {
		sequence := atomic.LoadInt64(c.Sequence)
		c.wsMutex.Lock()
		err := c.wsConn.WriteJSON(HeartbeatPayload{1, sequence})
		c.wsMutex.Unlock()

		if err != nil {
			c.Log.Errorf("Websocket: was unable to send heartbeat: %v", err)
			return
		} else if time.Now().Sub(c.LastHeartbeatAck) > c.HeartbeatInterval*c.MissingHeartbeatAcks {
			c.Log.Errorf("Websocket: did not receive a hearbeat acknowledgement for the last %v heartbeats", c.MissingHeartbeatAcks)
		}
		c.Log.Debug("Websocket: sent heartbeat to Discord")

		// Wait for next tick or quit
		select {
		case <-ticker.C:
			// continue loop
		case <-c.Connected:
			c.Log.Info("òwó")
			return
		}
	}
}

// Dispatch runs events through the plugin system
func (c *Client) Dispatch(event *Event) {

}

// Hello handles the initial Hello event
func (c *Client) Hello() error {

	// Retreive Hello message
	messageType, message, err := c.wsConn.ReadMessage()
	if err != nil {
		return err
	}

	// Retreive event out of message
	event, err := c.ResolveEvent(messageType, message)
	if err != nil {
		return err
	}
	if event.Opcode != opcodeHello {
		return fmt.Errorf("Websocket: expected opcode 10 Hello, received opcode %v %v instead", event.Opcode, event.Opcode.String())
	}

	// Resolve body
	var hello EventHello
	err = json.Unmarshal(event.RawData, &hello)
	if err != nil {
		return fmt.Errorf("Websocket: was unable to unmarshal Hello, %v", err)
	}
	c.HeartbeatInterval = hello.HeartbeatInterval * time.Millisecond
	c.LastHeartbeatAck = time.Now()

	return nil
}

// Login resumes the connecting or sends and identify payload
func (c *Client) Login() error {

	if c.SessionID == "" && atomic.LoadInt64(c.Sequence) == 0 {
		err := c.Identify()
		if err != nil {
			return err
		}
	} else {
		err := c.Resume()
		if err != nil {
			return err
		}
	}

	return nil
}

// Resume resumes a connection
func (c *Client) Resume() error {
	return nil
}

// Identify sends an identify payload, duh
func (c *Client) Identify() error {

	c.Log.Debug("Websocket: is authenticating with gateway...")

	// Build Identify data
	data := Identify{}
	data.Token = c.TOKEN
	data.Compress = c.Compress
	data.LargeThreshold = c.LargeThreshold
	data.Properties.OS = runtime.GOOS
	data.Properties.Browser = "Peach" + VERSION
	data.Properties.Device = "Peach" + VERSION
	data.Shard = [2]int{c.ShardID, c.ShardCount}

	// Create Payload
	payload := IdentifyPayload{2, data}

	if c.ShardID != 0 {
		queuetime, _ := time.ParseDuration(fmt.Sprintf("%vs", 6*c.ShardID))
		time.Sleep(queuetime)
	}
	// Send message
	c.wsMutex.Lock()
	err := c.wsConn.WriteJSON(payload)
	c.wsMutex.Unlock()
	c.Log.Debug("Websocket: sent Identify payload.")
	return err
}
