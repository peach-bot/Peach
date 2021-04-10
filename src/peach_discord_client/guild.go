package main

// Guild represents a discord guild
type Guild struct {
	ID                          string            `json:"id"`
	Name                        string            `json:"name"`
	Icon                        string            `json:"icon"`
	IconHash                    string            `json:"icon_hash"`
	Splash                      string            `json:"splash"`
	DiscoverySplash             string            `json:"discovery_splash"`
	IsOwner                     bool              `json:"owner,omitempty"`
	OwnerID                     string            `json:"owner_id"`
	Permissions                 string            `json:"permissions,omitempty"`
	Region                      string            `json:"region"`
	AFKChannelID                string            `json:"afk_channel_id"`
	AFKTimeout                  int               `json:"afk_timeout"`
	WidgetEnabled               bool              `json:"widget_enabled,omitempty"`
	WidgetChannelID             string            `json:"widget_channel_id,omitempty"`
	VerificationLevel           int               `json:"verification_level"`
	DefaultMessageNotifications int               `json:"default_message_notifications"`
	ExplicitContentFilter       int               `json:"explicit_content_filter"`
	Roles                       []*Role           `json:"roles"`
	Emojis                      []*Emoji          `json:"emojis"`
	Features                    []string          `json:"features"`
	MFALevel                    int               `json:"mfa_level"`
	ApplicationID               string            `json:"application_id"`
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
	MaxVideoChannelUsers        int               `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      int               `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    int               `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               WelcomeScreen     `json:"welcome_screen,omitempty"`
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

type WelcomeScreen struct {
	Description     string            `json:"description"`
	WelcomeChannels []*WelcomeChannel `json:"welcome_channels"`
}

type WelcomeChannel struct {
	ChannelID   string `json:"channel_id"`
	Description string `json:"description"`
	EmojiID     string `json:"emoji_id"`
	EmojiName   string `json:"emoji_name"`
}

type Ban struct {
	Reason string `json:"reason"`
	User   User   `json:"user"`
}

type Invite struct {
	Code                     string    `json:"code"`
	Guild                    Guild     `json:"guild,omitempty"`
	Channel                  Channel   `json:"channel"`
	Inviter                  User      `json:"inviter,omitempty"`
	TargetUser               User      `json:"target_user,omitempty"`
	TargetUserType           int       `json:"target_user_type,omitempty"`
	ApproximatePresenceCount int       `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   int       `json:"approximate_member_count,omitempty"`
	Uses                     int       `json:"uses,omitempty"`
	MaxUses                  int       `json:"max_uses,omitempty"`
	MaxAge                   int       `json:"max_age,omitempty"`
	Temporary                bool      `json:"temporary,omitempty"`
	CreatedAt                Timestamp `json:"created_at,omitempty"`
}

const (
	TargetUserTypeStream = iota + 1
)

type GuildPreview struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	Icon                     string   `json:"icon"`
	Splash                   string   `json:"splash"`
	DiscoverySplash          string   `json:"discovery_splash"`
	Emojis                   []*Emoji `json:"emojis"`
	Features                 []string `json:"features"`
	ApproximateMemberCount   int      `json:"approximate_member_count"`
	ApproximatePresenceCount int      `json:"approximate_presence_count"`
	Description              string   `json:"description"`
}

type Integration struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Type              int         `json:"type"`
	Enabled           bool        `json:"enabled"`
	Syncing           bool        `json:"syncing,omitempty"`
	RoleID            string      `json:"role_id,omitempty"`
	EnableEmoticons   bool        `json:"enable_emoticons,omitempty"`
	ExpireBehavior    int         `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int         `json:"expire_grace_period,omitempty"`
	User              User        `json:"user,omitempty"`
	Account           Account     `json:"account"`
	SyncedAt          Timestamp   `json:"synced_at,omitempty"`
	SubscriberCount   int         `json:"subscriber_count,omitempty"`
	Revoked           bool        `json:"revoked,omitempty"`
	Application       Application `json:"application,omitempty"`
}

type Widget struct {
	Enabled   bool   `json:"enabled"`
	ChannelID string `json:"channel_id"`
}

const (
	ExpireBehaviorRemoveRole int = iota
	ExpireBehaviorKick
)

// Guild features
const (
	GuildFeatureInviteSplash     = "INVITE_SPLASH"
	GuildFeatureVIPRegions       = "VIP_REGIONS"
	GuildFeatureVanityURL        = "VANITY_URL"
	GuildFeatureVerified         = "VERIFIED"
	GuildFeaturePartnered        = "PARTNERED"
	GuildFeaturePublic           = "PUBLIC"
	GuildFeatureCommerce         = "COMMERCE"
	GuildFeatureNews             = "NEWS"
	GuildFeatureDiscoverable     = "DISCOVERABLE"
	GuildFeatureFeaturable       = "FEATURABLE"
	GuildFeatureAnimatedIcon     = "ANIMATED_ICON"
	GuildFeatureBanner           = "BANNER"
	GuildFeatureWelcomeScreen    = "WELCOME_SCREEN_ENABLED"
	GuildFeatureVerificationGate = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeaturePreviewEnabled   = "PREVIEW_ENABLED"
)
