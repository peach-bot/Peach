package main

import (
	"encoding/json"
	"time"
)

// Event provides a basic initial struct for all websocket events.
type Event struct {
	OpCode   int             `json:"op"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
	RawData  json.RawMessage `json:"d"`
	// Struct contains one of the other types in this file.
	Struct interface{} `json:"-"`
}

// EventHello is the initial event sent by discord upon connection
type EventHello struct {
	HeartbeatInterval time.Duration `json:"hearbeat_interval"`
}
