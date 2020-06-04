package main

//go:generate go run cmd/eventresolvers/main.go

import (
	"time"
)

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

// EventMessageCreate is the data for a MessageCreate event.
type EventMessageCreate struct {
	*Message
}

// EventPresenceUpdate is the data for a PresenceUpdate event.
type EventPresenceUpdate struct {
	PresenceUpdate
}
