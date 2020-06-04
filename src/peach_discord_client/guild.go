package main

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

// UnavailableGuild represents an Offline Guild, or a Guild whose information has
// not been provided through Guild Create events during the Gateway connect.
type UnavailableGuild struct {
	ID          string `json:"id"`
	Unavailable bool   `json:"unavailable"`
}

// GuildMember represents a member of a discord guild
type GuildMember struct {
	User          User      `json:"user,omitempty"`
	Nickname      string    `json:"nick,omitempty"`
	Roles         []*string `json:"roles"`
	JoinedAt      string    `json:"joined_at"`
	BoostingSince string    `json:"premium_since,omitempty"`
	Deaf          bool      `json:"deaf"`
	Mute          bool      `json:"mute"`
	GuildID       string    `json:"guild_id,omitempty"`
}

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
