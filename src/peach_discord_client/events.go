package main

//go:generate go run cmd/eventresolvers/main.go

import (
	"time"
)

//
// CONNECTING AND RESUMING
//

// EventHello sent on connection to the websocket.
type EventHello struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval"`
}

// EventReady is dispatched when a client has completed the initial handshake with the gateway.
type EventReady struct {
	Version         int                 `json:"v"`
	User            User                `json:"user"`
	PrivateChannels []*Channel          `json:"private_channels"`
	Guilds          []*UnavailableGuild `json:"guilds"`
	SessionID       string              `json:"session_id"`
	Shard           *[2]int             `json:"shard,omitempty"`
}

// EventResumed is dispatched when a client has sent a resume payload to the gateway.
type EventResumed struct {
}

// EventReconnect is dispatched when a client should reconnect to the gateway
type EventReconnect struct {
}

//
// CHANNELS
//

// EventChannelCreate is sent when a new channel is created.
type EventChannelCreate struct {
	*Channel
}

// EventChannelUpdate is sent when a channel is updated.
type EventChannelUpdate struct {
	*Channel
}

// EventChannelDelete is sent when a channel relevant to the current user is deleted.
type EventChannelDelete struct {
	*Channel
}

// EventChannelPinsUpdate is sent when a message is pinned or unpinned in a text channel.
type EventChannelPinsUpdate struct {
	GuildID          string `json:"guild_id,omitempty"`
	ChannelID        string `json:"channel_id"`
	LastPinTimestamp string `json:"last_pin_timestamp"`
}

//
// GUILDS
//

// EventGuildCreate is sent when the user joins a new Guild or a Guild becomes available.
type EventGuildCreate struct {
	*Guild
}

// EventGuildUpdate is sent when a guild is updated.
type EventGuildUpdate struct {
	*Guild
}

// EventGuildDelete is sent when a guild becomes unavailable or the user leaves a guild.
type EventGuildDelete struct {
	*UnavailableGuild
}

// EventGuildBanAdd is sent when a user is banned from a guild.
type EventGuildBanAdd struct {
	GuildID string `json:"guild_id"`
	User    User   `json:"user"`
}

// EventGuildBanRemove is sent when a user is unbanned from a guild.
type EventGuildBanRemove struct {
	GuildID string `json:"guild_id"`
	User    User   `json:"user"`
}

// EventGuildEmojisUpdate is sent when a guild's emojis have been updated.
type EventGuildEmojisUpdate struct {
	GuildID string   `json:"guild_id"`
	Emojis  []*Emoji `json:"emojis"`
}

// EventGuildIntegrationsUpdate is sent when a guild integration is updated.
type EventGuildIntegrationsUpdate struct {
	GuildID string `json:"guild_id"`
}

// EventGuildMemberAdd is sent when a new user joins a guild.
type EventGuildMemberAdd struct {
	GuildMember
}

// EventGuildMemberRemove is sent when a user is removed or leaves a guild.
type EventGuildMemberRemove struct {
	GuildID string `json:"guild_id"`
	User    User   `json:"user"`
}

// EventGuildMemberUpdate is sent when a guild member is updated.
type EventGuildMemberUpdate struct {
	GuildID       string    `json:"guild_id"`
	Roles         []*string `json:"roles"`
	User          User      `json:"user"`
	Nickname      string    `json:"nick,omitempty"`
	BoostingSince string    `json:"premium_since,omitempty"`
}

// EventGuildRoleCreate is sent when a guild role is created.
type EventGuildRoleCreate struct {
	GuildID string `json:"guild_id"`
	Role    Role   `json:"role"`
}

// EventGuildRoleUpdate is sent when a guild role is updated.
type EventGuildRoleUpdate struct {
	GuildID string `json:"guild_id"`
	Role    Role   `json:"role"`
}

// EventGuildRoleDelete is sent when a guild role is deleted.
type EventGuildRoleDelete struct {
	GuildID string `json:"guild_id"`
	Role    string `json:"role"`
}

//
// INVITES
//

// EventInviteCreate is sent when a new invite to a channel is created.
type EventInviteCreate struct {
	ChannelID      string `json:"channel_id"`
	Code           string `json:"code"`
	CreatedAt      int    `json:"created_at"`
	GuildID        string `json:"guild_id,omitempty"`
	Inviter        User   `json:"inviter,omitempty"`
	MaxAge         int    `json:"max_age"`
	MaxUses        int    `json:"max_uses"`
	TargetUser     User   `json:"target_user,omitempty"`
	TargetUserType int    `json:"target_user_type,omitempty"`
	Temporary      bool   `json:"temporary"`
	Uses           int    `json:"uses"`
}

// EventInviteDelete is sent when an invite is deleted.
type EventInviteDelete struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
	Code      string `json:"code"`
}

//
// MESSAGES
//

// EventMessageCreate is the data for a MessageCreate event.
type EventMessageCreate struct {
	*Message
}

// EventTypingStart is sent when a user starts typing in a channel.
type EventTypingStart struct {
	ChannelID string      `json:"channel_id"`
	GuildID   string      `json:"guild_id,omitempty"`
	UserID    string      `json:"user_id"`
	Timestamp int         `json:"timestamp"`
	Member    GuildMember `json:"member,omitempty"`
}

//
// PRESENCE
//

// EventPresenceUpdate is the data for a PresenceUpdate event.
type EventPresenceUpdate struct {
	PresenceUpdate
}

//
// VOICE
//

//
// WEBHOOKS
//

// EventWebhooksUpdate is sent when a guild channel's webhook is created, updated, or deleted.
type EventWebhooksUpdate struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}
