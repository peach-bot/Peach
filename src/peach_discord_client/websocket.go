package main

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
		c.Connected = false
		c.log.WithField("Websocket closed:", text)
		fmt.Print("odslodas")
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
	c.Connected = true

	go c.Listen(c.wsConn)
	go c.Heartbeat(c.wsConn)

	c.log.Info("Websocket created")
	return nil
}

// Listen retrieves messages from the websocket
func (c *Client) Listen(wsConn *websocket.Conn) {
	c.log.Info("Websocket started listening")

	for {
		messageType, message, err := wsConn.ReadMessage()

		if err != nil {

		}

		c.ResolveEvent(messageType, message)

		if !c.Connected {
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

	return e, nil
}

// Heartbeat sends heartbeat payloads to discord to signal discord that the client is still alive
func (c *Client) Heartbeat(wsConn *websocket.Conn) {

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
	c.log.Info("Received opcode 10 Hello")

	// Resolve body
	var hello EventHello
	err = json.Unmarshal(event.RawData, &hello)
	if err != nil {
		return fmt.Errorf("Couldn't unmarshal Hello, %v", err)
	}
	c.heartbeatInterval = hello.HeartbeatInterval

	c.log.Info(hello.HeartbeatInterval)

	return nil
}

// Login resumes the connecting or sends and identify payload
func (c *Client) Login() error {
	err := c.Identify()
	if err != nil {
		return err
	}

	return nil
}

// Resume resumes a connection
func (c *Client) Resume() error {
	return nil
}

// Identify sends an identify payload, duh
func (c *Client) Identify() error {
	return nil
}
