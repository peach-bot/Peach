// This file containst structs for the client
// For structs representing discord objects see emoji.go, channel.go, guild.go, user.go, permissions.go, voice.go
// For websocket events see events.go
// For constants like close and opcodes see consts.go

package main

import (
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

// PresenceUpdate represents a discord presence update
type PresenceUpdate struct {
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
