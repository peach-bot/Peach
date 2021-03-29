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

	AddEventTypeHandlers()

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

	go c.Heartbeat()
	go c.Listen()

	c.Log.Info("Websocket: created")
	return nil
}

// Listen retrieves messages from the websocket
func (c *Client) Listen() {
	c.Log.Info("Websocket: started listening")

	for {
		// Read message from connection
		messageType, message, err := c.wsConn.ReadMessage()
		if err != nil {
			switch t := err.(type) {
			case *websocket.CloseError:
				c.Log.Error(t)
				c.Log.Debug("restarting websocket")
				c.Log.Debug("sent disconnect")
				if closecode(t.Code) == closecodeReconnect {
					c.Reconnect <- nil
					c.Connected <- nil
					return
				}
				c.Reconnect <- nil
				c.Connected <- nil
				return
			default:
				c.Log.Errorf("Websocket: was unable to read message: %v", err)
				return
			}
		} else {

			// Resolve event
			e, err := c.DecodeMessage(messageType, message, false)
			if err != nil {
				c.Log.Error(err)
			}
			go c.HandleEvent(e)
			if e.Opcode == opcodeReconnect {
				c.Reconnect <- nil
				c.Connected <- nil
				return
			}
		}

		select {
		case <-c.Connected:
			c.Log.Info("Websocket: stopped listening")
			return
		default:
		}
	}
}

// DecodeMessage decodes a websocket message, duh
func (c *Client) DecodeMessage(messageType int, message []byte, invalidSession bool) (*Event, error) {

	var reader io.Reader
	reader = bytes.NewBuffer(message)

	if messageType == websocket.BinaryMessage {
		z, _ := zlib.NewReader(reader)
		reader = z
	}
	var e *Event = new(Event)

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&e)
	if err != nil {
		err = fmt.Errorf("Websocket: Could not decode Event: %s", err)
		return nil, err
	}

	return e, nil
}

// HandleEvent resolves messages and handles the events included within
func (c *Client) HandleEvent(e *Event) {

	if e == nil {
		c.Log.Error("Websocket: Event is nil.")
	}

	if e.RawData == nil {
		c.Log.Error("Websocket: RawData is nil.")
	}

	// Store sequence
	atomic.StoreInt64(c.Sequence, e.Sequence)

	// Do opcode specific things
	if e.Opcode == opcodeHeartbeatACK {
		c.LastHeartbeatAck = time.Now()
		c.Log.Debug("Websocket: received heartbeat acknowledgement")
		return
	}

	if e.Opcode == opcodeInvalidSession {
		c.Log.Info("Websocket: invalid session")
		err := c.Identify()
		if err != nil {
			c.Log.Errorf("Error sending identify payload after receiving opcode inavlid session: %s", err)
		}
		return
	}

	if e.Opcode == opcodeDispatch {
		c.Log.Debugf("Websocket: received event %s from Discord", e.Type)
		eventtypehandler := eventTypeHandlers[e.Type]
		e.Struct = eventtypehandler.New()
		if err := json.Unmarshal(e.RawData, &e.Struct); err != nil {
			c.Log.Errorf("Websocket: Error unmarshalling %s event, %s", e.Type, err)
		}

		eventtypehandler.Handle(c, e.Struct)
	}

	if e.Opcode == opcodeReconnect {
		c.Connected <- nil
		c.wsConn.Close()
		c.Reconnect <- nil
	}
}

// Heartbeat sends heartbeat payloads to discord to signal discord that the client is still alive
func (c *Client) Heartbeat() {

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
			c.Log.Info("Websocket: Stopped heartbeat")
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
	event, err := c.DecodeMessage(messageType, message, false)
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

	c.Log.Info("Websocket: authenticating with gateway...")

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

	c.Log.Info("Websocket: resuming gateway connection...")

	// Build Resume payload
	data := Resume{}
	data.Sequence = atomic.LoadInt64(c.Sequence)
	data.SessionID = c.SessionID
	data.Token = c.TOKEN
	payload := ResumePayload{6, data}

	// Send message
	c.wsMutex.Lock()
	err := c.wsConn.WriteJSON(payload)
	c.wsMutex.Unlock()
	c.Log.Debug("Websocket: sent Resume payload")
	return err
}

// Identify sends an identify payload, duh
func (c *Client) Identify() error {

	c.Log.Info("Websocket: identifying...")

	// Build Identify payload

	data := Identify{}
	if c.Sharded {
		data := IdentifyWithShards{}
		data.Shard = [2]int{c.ShardID, c.ShardCount}
	}
	data.Token = c.TOKEN
	data.Compress = c.Compress
	data.LargeThreshold = c.LargeThreshold
	data.Properties.OS = runtime.GOOS
	data.Properties.Browser = "Peach" + VERSION
	data.Properties.Device = "Peach" + VERSION
	payload := IdentifyPayload{2, data}

	if c.Sharded {
		queuetime, _ := time.ParseDuration(fmt.Sprintf("%vs", 6*c.ShardID))
		time.Sleep(queuetime)
	}

	// Send message
	c.wsMutex.Lock()
	err := c.wsConn.WriteJSON(payload)
	c.wsMutex.Unlock()
	c.Log.Debug("Websocket: sent identify payload")
	return err
}
