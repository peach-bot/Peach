// This file containst structs for the client
// For structs representing discord objects see emoji.go, channel.go, guild.go, user.go, permissions.go, voice.go
// For websocket events see events.go
// For constants like close and opcodes see consts.go

package main

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Client represents connection to discord.
type Client struct {

	// Logger
	Log *logrus.Logger

	// Authentification
	TOKEN string

	// Settings
	Compress           bool
	LargeThreshold     int // total number of members where the gateway will stop sending offline members in the guild member list
	GuildSubscriptions bool
	Intents            int

	// Sharding
	ShardID    int
	ShardCount int

	// Gateway URL
	GatewayURL string

	// Shard Coordinator
	ShardCoordinatorURL string

	// Connected represents the clients connection status
	Connected chan interface{}

	// Session
	SessionID string
	Sequence  *int64

	// Heartbeat
	HeartbeatInterval    time.Duration // Interval in which client should sent heartbeats
	LastHeartbeatAck     time.Time     // Last time the client received a heartbeat acknowledgement
	MissingHeartbeatAcks time.Duration // Number of Acks that can be missed before reconnecting

	// Websocket Connection
	wsConn  *websocket.Conn
	wsMutex sync.Mutex
	sync.RWMutex

	// Snowflake node to generate snowflakes
	Snowflake snowflake.Node
}

// IdentifyPayload is used to create an identify message
type IdentifyPayload struct {
	Opcode int      `json:"op"`
	Data   Identify `json:"d"`
}

// Identify is used to trigger the initial handshake with the gateway.
type Identify struct {
	Token              string       `json:"token"`
	Compress           bool         `json:"compress,omitempty"`
	LargeThreshold     int          `json:"large_threshold,omitemtpy"`
	Shard              [2]int       `json:"shard,omitempty"`
	Presence           UpdateStatus `json:"presence,omitempty"`
	GuildSubscriptions bool         `json:"guild_subscriptions,omitempty"`
	Intents            int          `json:"intents,omitempty"`
	Properties         struct {
		OS      string `json:"$os"`
		Browser string `json:"$browser"`
		Device  string `json:"$device"`
	} `json:"properties"`
}

// ResumePayload is used to create a resume message
type ResumePayload struct {
	Opcode int    `json:"op"`
	Data   Resume `json:"d"`
}

// Resume is used to resume a connection
type Resume struct {
}

// HeartbeatPayload is used to create a heartbeat message
type HeartbeatPayload struct {
	Opcode int   `json:"op"`
	Data   int64 `json:"d"`
}

// ShardCoordinatorResponse is used to unmarshal the shard coordinator response
type ShardCoordinatorResponse struct {
	TotalShards int    `json:"total_shards"`
	ShardID     int    `json:"assigned_shard"`
	GatewayURL  string `json:"gatewayurl"`
}

// UpdateStatus is sent by the client to indicate a presence or status update.
type UpdateStatus struct {
	Since    int      `json:"since"`  // unix time in ms or null
	Activity Activity `json:"game"`   // the clients new activity or null
	Status   string   `json:"status"` // the clients new status
	AFK      bool     `json:"afk"`    // whether the client is afk or not
}

// Activity represence a discord status activity
type Activity struct {
	Name      string `json:"name"`
	Type      int    `json:"type"`
	CreatedAt int    `json:"created_at"`
}

// ClientStatus represents a user's active sessions
type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Web     string `json:"web,omitempty"`
}

// PresenceUpdate represents an update to a user's current state on a guild
type PresenceUpdate struct {
	User         User         `json:"user"`
	Roles        []string     `json:"roles"`
	Game         Activity     `json:"game,omitempty"`
	GuildID      string       `json:"guild_id"`
	Status       string       `json:"status"`
	Activities   []*Activity  `json:"activities"`
	ClientStatus ClientStatus `json:"client_status"`
	NitroSince   string       `json:"premium_since,omitempty"`
	Nickname     string       `json:"nick,omitempty"`
}

// Event provides a basic initial struct for all websocket events.
type Event struct {
	Opcode   opcode          `json:"op"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
	RawData  json.RawMessage `json:"d"`
	// Struct contains one of the other Events
	Struct interface{} `json:"-"`
}

// EventInvalidSession is sent to indicate that the session could not be initialized, resumed or was invalidated.
type EventInvalidSession struct {
	Opcode   opcode `json:"op"`
	Sequence int64  `json:"s"`
	Type     string `json:"t"`
	RawData  bool   `json:"d"`
}
