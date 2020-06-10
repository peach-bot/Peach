package main

// VoiceState represents a discord voice state
type VoiceState struct {
	GuildID    string      `json:"guild_id,omitempty"`
	ChannelID  string      `json:"channel_id,omitempty"`
	UserID     string      `json:"user_id"`
	Member     GuildMember `json:"member,omitempty"`
	SessionID  string      `json:"session_id"`
	Deaf       bool        `json:"deaf"`
	Mute       bool        `json:"mute"`
	SelfDeaf   bool        `json:"self_deaf"`
	SelfMute   bool        `json:"self_mute"`
	SelfStream bool        `json:"self_stream,omitempty"`
	Suppress   bool        `json:"suppress"`
}

// VoiceRegion represents a voicechat region
type VoiceRegion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	VIP        bool   `json:"vip"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}
