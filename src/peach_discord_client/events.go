package main

import (
	"encoding/json"
	"time"
)

// Event provides a basic initial struct for all websocket events.
type Event struct {
	Opcode   opcode          `json:"op"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
	RawData  json.RawMessage `json:"d"`
	// Struct contains one of the other types in this file.
	Struct interface{} `json:"-"`
}

// EventHello is the initial event sent by discord upon connection
type EventHello struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

// EventReady resembles a Ready event
type EventReady struct {
	Version         int        `json:"v"`
	User            User       `json:"user"`
	PrivateChannels []*Channel `json:"private_channels"`
	Guilds          []*Guild   `json:"guilds"`
	SessionID       string     `json:"session_id"`
	Shard           *[2]int    `json:"shard,omitempty"`
}
