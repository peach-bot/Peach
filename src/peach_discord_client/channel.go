package main

// Channel represents a discord channel
type Channel struct {
	ID                   string       `json:"id"`
	Type                 int          `json:"type"`
	GuildID              string       `json:"guild_id,omitempty"`
	Position             int          `json:"position,omitempty"`
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"`
	Name                 string       `json:"string,omitempty"`
	Topic                string       `json:"topic,omitempty"`
	NSFW                 bool         `json:"nsfw,omitempty"`
	LastMessageID        string       `json:"last_message_id,omitempty"`
	Bitrate              int          `json:"bitrate,omitempty"`
	UserLimit            int          `json:"user_limit,omitempty"`
	RateLimitPerUser     int          `json:"rate_limit_per_user,omitempty"`
	Recipients           []*User      `json:"recipients,omitempty"`
	Icon                 string       `json:"icon,omitempty"`
	OwnerID              string       `json:"owner_id,omitempty"`
	ApplicationID        string       `json:"application_id,omitempty"`
	ParentID             string       `json:"partent_id,omitempty"`
	LastPinTimestamp     string       `json:"last_pin_timestamp,omitempty"`
	RTCRegion            string       `json:"rtc_region,omitempty"`
	VideoQualityMode     int          `json:"video_quality_mode,omitempty"`
}

// Overwrite represents an explicit permission overwrite for members or roles
type Overwrite struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}

// Message represents a discord message
type Message struct {
	ID               string             `json:"id"`
	ChannelID        string             `json:"channel_id"`
	GuildID          string             `json:"guild_id,omitempty"`
	Author           User               `json:"author"`
	Member           GuildMember        `json:"member,omitempty"`
	Content          string             `json:"content"`
	Timestamp        string             `json:"timestamp"`
	EditedTimestamp  string             `json:"edited_timestamp,omitempty"`
	TTS              bool               `json:"tts"`
	MentionEveryone  bool               `json:"mention_everyone"`
	Mentions         []*User            `json:"mentions"`
	MentionRoles     []*string          `json:"mention_roles"`
	MentionChannels  []*ChannelMention  `json:"mention_channels,omitempty"`
	Attachments      []*Attachment      `json:"attachments"`
	Embeds           []*Embed           `json:"embeds"`
	Reactions        []*Reaction        `json:"reactions,omitempty"`
	Pinned           bool               `json:"pinned"`
	WebhookID        string             `json:"webhook_id,omitempty"`
	Type             messagetype        `json:"type"`
	Activity         MessageActivity    `json:"activity,omitempty"`
	Application      MessageApplication `json:"application,omitempty"`
	MessageReference MessageReference   `json:"message_reference,omitempty"`
	Flags            int                `json:"flags,omitempty"`
}

const (
	MessageFlagCrossposted = 1 << iota
	MessageFlagIsCrossposted
	MessageFlagSuppressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
	MessageFlagEphemeral
	MessageFlagLoading
)

func (m *Message) Delete(c *Client) error {
	return c.DeleteMessage(m.ChannelID, m.ID)
}

func (m *Message) Edit(c *Client, args EditMessageArgs) (*Message, error) {
	return c.EditMessage(m.ChannelID, m.ID, args)
}

// Attachment represents a Discord message's attachment
type Attachment struct {
	ID       string `json:"string"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

// Reaction represents reactions to a Discord message
type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

// ChannelMention represents a mentioned channel in a Discord message
type ChannelMention struct {
	ID      string `json:"id"`
	GuildID string `json:"guild_id"`
	Type    int    `json:"type"`
	Name    string `json:"name"`
}

// MessageReference represents the reference data sent with crossposted messages
type MessageReference struct {
	MessageID string `json:"message_id,omitempty"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
}

// MessageActivity is sent with Rich Presence-related chat embeds, for example a party invite
type MessageActivity struct {
	Type    int    `json:"type"`
	PartyID string `json:"party_id"`
}

// MessageApplication is sent with Rich Presence-related chat embeds, for example if a Fortnite party invite has been sent this would represent Fortnite
type MessageApplication struct {
	ID          string `json:"id"`
	CoverImage  string `json:"cover_image,omitempty"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
}

// Embed represents a Discord embed, and that's a fact
type Embed struct {
	Title       string         `json:"title,omitempty"`
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitepmty"`
	Timestamp   string         `json:"timestamp,omitempty"`
	Color       int            `json:"color"`
	Footer      EmbedFooter    `json:"footer,omitempty"`
	Image       EmbedImage     `json:"image,omitempty"`
	Thumbnail   EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       EmbedVideo     `json:"video,omitempty"`
	Provider    EmbedProvider  `json:"provider,omitempty"`
	Author      EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField  `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedField represents a Discord embed field
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type AllowedMentions struct {
	MentionTypes       []MentionType `json:"parse"`
	RoleIDs            []string      `json:"roles"`
	UserIDs            []string      `json:"users"`
	MentionRepliedUser bool          `json:"replied_user"`
}

type MentionType string

const (
	MentionRoles    MentionType = "roles"
	MentionUsers                = "users"
	MentionEveryone             = "everyone"
)
