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
	c.wsConn.SetCloseHandler(func(code int, text string) error {
		c.Connected = nil
		c.Log.WithField("Websocket closed:", code)
		return nil
	})

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

	go c.Listen(c.wsConn)
	go c.Heartbeat(c.wsConn)

	c.Log.Info("Websocket created")
	return nil
}

// Listen retrieves messages from the websocket
func (c *Client) Listen(wsConn *websocket.Conn) {
	c.Log.Info("Websocket started listening")

	for {
		// Read message from connection
		messageType, message, err := wsConn.ReadMessage()
		if err != nil {

		}

		// Resolve event
		_, err = c.ResolveEvent(messageType, message)
		if err != nil {
			c.Log.Error(err)
		}

		select {
		case <-c.Connected:
			break
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
		return e, err
	}

	// Store sequence
	atomic.StoreInt64(c.Sequence, e.Sequence)

	// Do opcode specific things
	if e.OpCode == opCodeHeartbeatACK {
		c.LastHeartbeatAck = time.Now()
		c.Log.Info("Websocket received opcode 11 HeartbeatACK.")
	} else if e.OpCode == opCodeDispatch {
		c.Log.Infof("Websocket received opcode 0 Dispatch with event %s from Discord.", e.Type)
	} else if e.OpCode == opCodeHello {
		c.Log.Info("Received opcode 10 Hello from Discord.")
	} else {
		c.Log.Infof("Websocket received opcode %v.", e.OpCode)
	}

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

		if err != nil || time.Now().Sub(c.LastHeartbeatAck) > c.HeartbeatInterval*c.MissingHeartbeatAcks {

		}
		c.Log.Info("Sent heartbeat to Discord")

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
	if event.OpCode != opCodeHello {
		return fmt.Errorf("Expected opcode 10 Hello, received opcode %v intead", event.OpCode)
	}

	// Resolve body
	var hello EventHello
	err = json.Unmarshal(event.RawData, &hello)
	if err != nil {
		return fmt.Errorf("Couldn't unmarshal Hello, %v", err)
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

	c.Log.Info("Sending Identify payload...")

	// Build Identify data
	data := Identify{}
	data.Token = c.TOKEN
	data.Compress = c.Compress
	data.LargeThreshold = c.LargeThreshold
	data.Properties.OS = runtime.GOOS
	data.Properties.Browser = "Peach" + VERSION
	data.Properties.Device = "Peach" + VERSION

	// Create Payload
	payload := IdentifyPayload{2, data}

	// Send message
	c.wsMutex.Lock()
	err := c.wsConn.WriteJSON(payload)
	c.wsMutex.Unlock()
	c.Log.Info("Sent payload uwu")
	return err
}
