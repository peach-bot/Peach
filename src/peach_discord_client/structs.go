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
	ShardCoordinator    ShardCoordinatorResponse

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
	OpCode int      `json:"op"`
	Data   Identify `json:"d"`
}

// Identify is used to trigger the initial handshake with the gateway.
type Identify struct {
	Token              string       `json:"token"`
	Compress           bool         `json:"compress,omitempty"`
	LargeThreshold     int          `json:"large_threshold,omitemtpy"`
	Shard              *[2]int      `json:"shard,omitempty"`
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
	OpCode int    `json:"op"`
	Data   Resume `json:"d"`
}

// Resume is used to resume a connection
type Resume struct {
}

// HeartbeatPayload is used to create a heartbeat message
type HeartbeatPayload struct {
	OpCode int   `json:"op"`
	Data   int64 `json:"d"`
}

// ShardCoordinatorResponse is used to unmarshal the shard coordinator response
type ShardCoordinatorResponse struct {
	TotalShards int  `json:"total_shards"`
	ShardID     int  `json:"assigned_shard"`
	APIShardID  int  `json:"api_shardid"`
	IsServer    bool `json:"is_server"`
}

// User represents a discord user
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot,omitempty"`
	System        bool   `json:"system,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Language      string `json:"locale,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Email         string `json:"email,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	NitroType     int    `json:"premium_type,omitempty"`
}

// Guild represents a discord guild
type Guild struct {
	ID                          string            `json:"id"`
	Name                        string            `json:"name"`
	Icon                        string            `json:"icon"`
	Splash                      string            `json:"splash"`
	DiscoverySplash             string            `json:"discovery_splash"`
	IsOwner                     bool              `json:"owner,omitempty"`
	OwnerID                     string            `json:"owner_id"`
	Permissions                 int               `json:"permissions,omitempty"`
	Region                      string            `json:"region"`
	AFKChannelID                string            `json:"afk_channel_id"`
	AFKTimeout                  int               `json:"afk_timeout"`
	EmbedEnabled                bool              `json:"embed_enabled,omitempty"`
	EmbedChannelID              string            `json:"embed_channel_id,omitempty"`
	VerificationLevel           int               `json:"verification_level"`
	DefaultMessageNotifications int               `json:"default_message_notifications"`
	ExplicitContentFilter       int               `json:"explicit_content_filter"`
	Roles                       []*Role           `json:"roles"`
	Emojis                      []*Emoji          `json:"emojis"`
	Features                    []string          `json:"features"`
	MFALevel                    int               `json:"mfa_level"`
	ApplicationID               string            `json:"application_id"`
	WidgetEnabled               bool              `json:"widget_enabled,omitempty"`
	WidgetChannelID             string            `json:"widget_channel_id,omitempty"`
	SystemChannelID             string            `json:"system_channel_id"`
	SystemChannelFlags          int               `json:"system_channel_flags"`
	RulesChannelID              string            `json:"rules_channel_id"`
	JoinedAt                    string            `json:"joined_at,omitempty"`
	Large                       bool              `json:"large,omitempty"`
	Unavailable                 bool              `json:"unavailable,omitempty"`
	MemberCount                 int               `json:"member_count,omitempty"`
	VoiceStates                 []*VoiceState     `json:"voice_states,omitempty"`
	Members                     []*GuildMember    `json:"members,omitempty"`
	Channels                    []*Channel        `json:"channels,omitempty"`
	Presences                   []*PresenceUpdate `json:"presences,omitempty"`
	MaxPresences                int               `json:"max_presences,omitempty"`
	MaxMembers                  int               `json:"max_members,omitempty"`
	VanityURLCode               string            `json:"vanity_url_code"`
	Description                 string            `json:"description,omitempty"`
	Banner                      string            `json:"banner,omitempty"`
	BoostLevel                  int               `json:"premium_tier"`
	Boosts                      int               `json:"premium_subscription_count,omitempty"`
	PreferredLanguage           string            `json:"preferred_locale"`
	PublicUpdatesChannelID      string            `json:"public_updates_channel_id"`
}

// Channel represents a discord channel
type Channel struct {
}

// Emoji represents a discord emoji
type Emoji struct {
}

// PresenceUpdate represents a discord presence update
type PresenceUpdate struct {
}

// GuildMember represents a member of a discord guild
type GuildMember struct {
}

// VoiceState represents a discord voice state
type VoiceState struct {
}

// Role represents a discord guild role
type Role struct {
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

// Gateway opcodes, denote payload type, see https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
const (
	opCodeDispatch            = 0  // Receive      | An event was dispatched.
	opCodeHeartbeat           = 1  // Send/Receive | Fired periodically by the client to keep the connection alive.
	opCodeIdentify            = 2  // Send         | Starts a new session during the initial handshake.
	opCodePresenceUpdate      = 3  // Send         | Update the client's presence.
	opCodeVoiceStateUpdate    = 4  // Send         | Used to join/leave or move between voice channels.
	opCodeResume              = 6  // Send         | Resume a previous session that was disconnected.
	opCodeReconnect           = 7  // Receive      | You must reconnect with a new session immediately.
	opCodeRequestGuildMembers = 8  // Send         | Request information about offline guild members in a large guild.
	opCodeInvalidSession      = 9  // Receive      | The session has been invalidated. You should reconnect and identify/resume accordingly.
	opCodeHello               = 10 // Receive      | Sent immediately after connecting, contains the heartbeat_interval to use.
	opCodeHeartbeatACK        = 11 // Receive      | Sent in response to receiving a heartbeat to acknowledge that it has been received.
)

// Gateway Close Event Codes, denote reason for gateway closure, see https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
const (
	closeCodeUnknownError         = 4000 // Not sure what went wrong. Try reconnecting.
	closeCodeUnknownOpCode        = 4001 // Sent invalid opcode or invalid payload for opcode.
	closeCodeDecodeError          = 4002 // Sent invalid payload.
	closeCodeNotAuthenticated     = 4003 // Sent payload prior to identifying.
	closeCodeAuthenticationFailed = 4004 // Account token in identify payload is incorrect.
	closeCodeAlreadyAuthenticated = 4005 // Sent more than one identify payload.
	closeCodeInvalidSquence       = 4007 // Sent invalid sequence when resuming.
	closeCodeRateLimited          = 4008 // Sending payloads to quickly.
	closeCodeSessionTimedOut      = 4009 // Session timed out. Reconnect or start new session.
	closeCodeInvalidShard         = 4010 // Sent invalid shard in identify payload.
	closeCodeShardingRequired     = 4011 // Sharding required because bot is in too many guilds.
	closeCodeInvalidAPIVersion    = 4012 // Sent an invalid gateway version.
	closeCodeInvalidIntents       = 4013 // Sent invalid gateway intent.
	closeCodeDisallowedIntents    = 4014 // Sent intent the account isn't eligible for.
)

// Guild features
const (
	GuildFeatureInviteSplash   = "INVITE_SPLASH"
	GuildFeatureVIPRegions     = "VIP_REGIONS"
	GuildFeatureVanityURL      = "VANITY_URL"
	GuildFeatureVerified       = "VERIFIED"
	GuildFeaturePartnered      = "PARTNERED"
	GuildFeaturePublic         = "PUBLIC"
	GuildFeatureCommerce       = "COMMERCE"
	GuildFeatureNews           = "NEWS"
	GuildFeatureDiscoverable   = "DISCOVERABLE"
	GuildFeatureFeaturable     = "FEATURABLE"
	GuildFeatureAnimatedIcon   = "ANIMATED_ICON"
	GuildFeatureBanner         = "BANNER"
	GuildFeaturePublicDisabled = "PUBLIC_DISABLED"
)
