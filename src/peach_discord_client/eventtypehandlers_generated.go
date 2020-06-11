package main

// Event types used to match values sent by Discord
const (
	channelCreateEventType              = "CHANNEL_CREATE"
	channelDeleteEventType              = "CHANNEL_DELETE"
	channelPinsUpdateEventType          = "CHANNEL_PINS_UPDATE"
	channelUpdateEventType              = "CHANNEL_UPDATE"
	guildBanAddEventType                = "GUILD_BAN_ADD"
	guildBanRemoveEventType             = "GUILD_BAN_REMOVE"
	guildCreateEventType                = "GUILD_CREATE"
	guildDeleteEventType                = "GUILD_DELETE"
	guildEmojisUpdateEventType          = "GUILD_EMOJIS_UPDATE"
	guildIntegrationsUpdateEventType    = "GUILD_INTEGRATIONS_UPDATE"
	guildMemberAddEventType             = "GUILD_MEMBER_ADD"
	guildMemberRemoveEventType          = "GUILD_MEMBER_REMOVE"
	guildMemberUpdateEventType          = "GUILD_MEMBER_UPDATE"
	guildRoleCreateEventType            = "GUILD_ROLE_CREATE"
	guildRoleDeleteEventType            = "GUILD_ROLE_DELETE"
	guildRoleUpdateEventType            = "GUILD_ROLE_UPDATE"
	guildUpdateEventType                = "GUILD_UPDATE"
	helloEventType                      = "HELLO"
	inviteCreateEventType               = "INVITE_CREATE"
	inviteDeleteEventType               = "INVITE_DELETE"
	messageCreateEventType              = "MESSAGE_CREATE"
	messageDeleteEventType              = "MESSAGE_DELETE"
	messageDeleteBulkEventType          = "MESSAGE_DELETE_BULK"
	messageReactionAddEventType         = "MESSAGE_REACTION_ADD"
	messageReactionRemoveEventType      = "MESSAGE_REACTION_REMOVE"
	messageReactionRemoveAllEventType   = "MESSAGE_REACTION_REMOVE_ALL"
	messageReactionRemoveEmojiEventType = "MESSAGE_REACTION_REMOVE_EMOJI"
	messageUpdateEventType              = "MESSAGE_UPDATE"
	presenceUpdateEventType             = "PRESENCE_UPDATE"
	readyEventType                      = "READY"
	reconnectEventType                  = "RECONNECT"
	resumedEventType                    = "RESUMED"
	typingStartEventType                = "TYPING_START"
	userUpdateEventType                 = "USER_UPDATE"
	voiceServerUpdateEventType          = "VOICE_SERVER_UPDATE"
	voiceStateUpdateEventType           = "VOICE_STATE_UPDATE"
	webhooksUpdateEventType             = "WEBHOOKS_UPDATE"
)

// channelCreateEventTypeHandler is an event handler for ChannelCreate events.
type channelCreateEventTypeHandler func(*Client, *EventChannelCreate)

// Type returns the event type for ChannelCreate events.
func (eventTypeHandler channelCreateEventTypeHandler) Type() string {
	return channelCreateEventType
}

// New returns a new instance of ChannelCreate.
func (eventTypeHandler channelCreateEventTypeHandler) New() interface{} {
	return &EventChannelCreate{}
}

// Handle is the handler for ChannelCreate events.
func (eventTypeHandler channelCreateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventChannelCreate)
	err := c.onChannelCreate(e)
	return err
}

// channelDeleteEventTypeHandler is an event handler for ChannelDelete events.
type channelDeleteEventTypeHandler func(*Client, *EventChannelDelete)

// Type returns the event type for ChannelDelete events.
func (eventTypeHandler channelDeleteEventTypeHandler) Type() string {
	return channelDeleteEventType
}

// New returns a new instance of ChannelDelete.
func (eventTypeHandler channelDeleteEventTypeHandler) New() interface{} {
	return &EventChannelDelete{}
}

// Handle is the handler for ChannelDelete events.
func (eventTypeHandler channelDeleteEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventChannelDelete)
	err := c.onChannelDelete(e)
	return err
}

// channelPinsUpdateEventTypeHandler is an event handler for ChannelPinsUpdate events.
type channelPinsUpdateEventTypeHandler func(*Client, *EventChannelPinsUpdate)

// Type returns the event type for ChannelPinsUpdate events.
func (eventTypeHandler channelPinsUpdateEventTypeHandler) Type() string {
	return channelPinsUpdateEventType
}

// New returns a new instance of ChannelPinsUpdate.
func (eventTypeHandler channelPinsUpdateEventTypeHandler) New() interface{} {
	return &EventChannelPinsUpdate{}
}

// Handle is the handler for ChannelPinsUpdate events.
func (eventTypeHandler channelPinsUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventChannelPinsUpdate)
	err := c.onChannelPinsUpdate(e)
	return err
}

// channelUpdateEventTypeHandler is an event handler for ChannelUpdate events.
type channelUpdateEventTypeHandler func(*Client, *EventChannelUpdate)

// Type returns the event type for ChannelUpdate events.
func (eventTypeHandler channelUpdateEventTypeHandler) Type() string {
	return channelUpdateEventType
}

// New returns a new instance of ChannelUpdate.
func (eventTypeHandler channelUpdateEventTypeHandler) New() interface{} {
	return &EventChannelUpdate{}
}

// Handle is the handler for ChannelUpdate events.
func (eventTypeHandler channelUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventChannelUpdate)
	err := c.onChannelUpdate(e)
	return err
}

// guildBanAddEventTypeHandler is an event handler for GuildBanAdd events.
type guildBanAddEventTypeHandler func(*Client, *EventGuildBanAdd)

// Type returns the event type for GuildBanAdd events.
func (eventTypeHandler guildBanAddEventTypeHandler) Type() string {
	return guildBanAddEventType
}

// New returns a new instance of GuildBanAdd.
func (eventTypeHandler guildBanAddEventTypeHandler) New() interface{} {
	return &EventGuildBanAdd{}
}

// Handle is the handler for GuildBanAdd events.
func (eventTypeHandler guildBanAddEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildBanAdd)
	err := c.onGuildBanAdd(e)
	return err
}

// guildBanRemoveEventTypeHandler is an event handler for GuildBanRemove events.
type guildBanRemoveEventTypeHandler func(*Client, *EventGuildBanRemove)

// Type returns the event type for GuildBanRemove events.
func (eventTypeHandler guildBanRemoveEventTypeHandler) Type() string {
	return guildBanRemoveEventType
}

// New returns a new instance of GuildBanRemove.
func (eventTypeHandler guildBanRemoveEventTypeHandler) New() interface{} {
	return &EventGuildBanRemove{}
}

// Handle is the handler for GuildBanRemove events.
func (eventTypeHandler guildBanRemoveEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildBanRemove)
	err := c.onGuildBanRemove(e)
	return err
}

// guildCreateEventTypeHandler is an event handler for GuildCreate events.
type guildCreateEventTypeHandler func(*Client, *EventGuildCreate)

// Type returns the event type for GuildCreate events.
func (eventTypeHandler guildCreateEventTypeHandler) Type() string {
	return guildCreateEventType
}

// New returns a new instance of GuildCreate.
func (eventTypeHandler guildCreateEventTypeHandler) New() interface{} {
	return &EventGuildCreate{}
}

// Handle is the handler for GuildCreate events.
func (eventTypeHandler guildCreateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildCreate)
	err := c.onGuildCreate(e)
	return err
}

// guildDeleteEventTypeHandler is an event handler for GuildDelete events.
type guildDeleteEventTypeHandler func(*Client, *EventGuildDelete)

// Type returns the event type for GuildDelete events.
func (eventTypeHandler guildDeleteEventTypeHandler) Type() string {
	return guildDeleteEventType
}

// New returns a new instance of GuildDelete.
func (eventTypeHandler guildDeleteEventTypeHandler) New() interface{} {
	return &EventGuildDelete{}
}

// Handle is the handler for GuildDelete events.
func (eventTypeHandler guildDeleteEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildDelete)
	err := c.onGuildDelete(e)
	return err
}

// guildEmojisUpdateEventTypeHandler is an event handler for GuildEmojisUpdate events.
type guildEmojisUpdateEventTypeHandler func(*Client, *EventGuildEmojisUpdate)

// Type returns the event type for GuildEmojisUpdate events.
func (eventTypeHandler guildEmojisUpdateEventTypeHandler) Type() string {
	return guildEmojisUpdateEventType
}

// New returns a new instance of GuildEmojisUpdate.
func (eventTypeHandler guildEmojisUpdateEventTypeHandler) New() interface{} {
	return &EventGuildEmojisUpdate{}
}

// Handle is the handler for GuildEmojisUpdate events.
func (eventTypeHandler guildEmojisUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildEmojisUpdate)
	err := c.onGuildEmojisUpdate(e)
	return err
}

// guildIntegrationsUpdateEventTypeHandler is an event handler for GuildIntegrationsUpdate events.
type guildIntegrationsUpdateEventTypeHandler func(*Client, *EventGuildIntegrationsUpdate)

// Type returns the event type for GuildIntegrationsUpdate events.
func (eventTypeHandler guildIntegrationsUpdateEventTypeHandler) Type() string {
	return guildIntegrationsUpdateEventType
}

// New returns a new instance of GuildIntegrationsUpdate.
func (eventTypeHandler guildIntegrationsUpdateEventTypeHandler) New() interface{} {
	return &EventGuildIntegrationsUpdate{}
}

// Handle is the handler for GuildIntegrationsUpdate events.
func (eventTypeHandler guildIntegrationsUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildIntegrationsUpdate)
	err := c.onGuildIntegrationsUpdate(e)
	return err
}

// guildMemberAddEventTypeHandler is an event handler for GuildMemberAdd events.
type guildMemberAddEventTypeHandler func(*Client, *EventGuildMemberAdd)

// Type returns the event type for GuildMemberAdd events.
func (eventTypeHandler guildMemberAddEventTypeHandler) Type() string {
	return guildMemberAddEventType
}

// New returns a new instance of GuildMemberAdd.
func (eventTypeHandler guildMemberAddEventTypeHandler) New() interface{} {
	return &EventGuildMemberAdd{}
}

// Handle is the handler for GuildMemberAdd events.
func (eventTypeHandler guildMemberAddEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildMemberAdd)
	err := c.onGuildMemberAdd(e)
	return err
}

// guildMemberRemoveEventTypeHandler is an event handler for GuildMemberRemove events.
type guildMemberRemoveEventTypeHandler func(*Client, *EventGuildMemberRemove)

// Type returns the event type for GuildMemberRemove events.
func (eventTypeHandler guildMemberRemoveEventTypeHandler) Type() string {
	return guildMemberRemoveEventType
}

// New returns a new instance of GuildMemberRemove.
func (eventTypeHandler guildMemberRemoveEventTypeHandler) New() interface{} {
	return &EventGuildMemberRemove{}
}

// Handle is the handler for GuildMemberRemove events.
func (eventTypeHandler guildMemberRemoveEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildMemberRemove)
	err := c.onGuildMemberRemove(e)
	return err
}

// guildMemberUpdateEventTypeHandler is an event handler for GuildMemberUpdate events.
type guildMemberUpdateEventTypeHandler func(*Client, *EventGuildMemberUpdate)

// Type returns the event type for GuildMemberUpdate events.
func (eventTypeHandler guildMemberUpdateEventTypeHandler) Type() string {
	return guildMemberUpdateEventType
}

// New returns a new instance of GuildMemberUpdate.
func (eventTypeHandler guildMemberUpdateEventTypeHandler) New() interface{} {
	return &EventGuildMemberUpdate{}
}

// Handle is the handler for GuildMemberUpdate events.
func (eventTypeHandler guildMemberUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildMemberUpdate)
	err := c.onGuildMemberUpdate(e)
	return err
}

// guildRoleCreateEventTypeHandler is an event handler for GuildRoleCreate events.
type guildRoleCreateEventTypeHandler func(*Client, *EventGuildRoleCreate)

// Type returns the event type for GuildRoleCreate events.
func (eventTypeHandler guildRoleCreateEventTypeHandler) Type() string {
	return guildRoleCreateEventType
}

// New returns a new instance of GuildRoleCreate.
func (eventTypeHandler guildRoleCreateEventTypeHandler) New() interface{} {
	return &EventGuildRoleCreate{}
}

// Handle is the handler for GuildRoleCreate events.
func (eventTypeHandler guildRoleCreateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildRoleCreate)
	err := c.onGuildRoleCreate(e)
	return err
}

// guildRoleDeleteEventTypeHandler is an event handler for GuildRoleDelete events.
type guildRoleDeleteEventTypeHandler func(*Client, *EventGuildRoleDelete)

// Type returns the event type for GuildRoleDelete events.
func (eventTypeHandler guildRoleDeleteEventTypeHandler) Type() string {
	return guildRoleDeleteEventType
}

// New returns a new instance of GuildRoleDelete.
func (eventTypeHandler guildRoleDeleteEventTypeHandler) New() interface{} {
	return &EventGuildRoleDelete{}
}

// Handle is the handler for GuildRoleDelete events.
func (eventTypeHandler guildRoleDeleteEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildRoleDelete)
	err := c.onGuildRoleDelete(e)
	return err
}

// guildRoleUpdateEventTypeHandler is an event handler for GuildRoleUpdate events.
type guildRoleUpdateEventTypeHandler func(*Client, *EventGuildRoleUpdate)

// Type returns the event type for GuildRoleUpdate events.
func (eventTypeHandler guildRoleUpdateEventTypeHandler) Type() string {
	return guildRoleUpdateEventType
}

// New returns a new instance of GuildRoleUpdate.
func (eventTypeHandler guildRoleUpdateEventTypeHandler) New() interface{} {
	return &EventGuildRoleUpdate{}
}

// Handle is the handler for GuildRoleUpdate events.
func (eventTypeHandler guildRoleUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildRoleUpdate)
	err := c.onGuildRoleUpdate(e)
	return err
}

// guildUpdateEventTypeHandler is an event handler for GuildUpdate events.
type guildUpdateEventTypeHandler func(*Client, *EventGuildUpdate)

// Type returns the event type for GuildUpdate events.
func (eventTypeHandler guildUpdateEventTypeHandler) Type() string {
	return guildUpdateEventType
}

// New returns a new instance of GuildUpdate.
func (eventTypeHandler guildUpdateEventTypeHandler) New() interface{} {
	return &EventGuildUpdate{}
}

// Handle is the handler for GuildUpdate events.
func (eventTypeHandler guildUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventGuildUpdate)
	err := c.onGuildUpdate(e)
	return err
}

// helloEventTypeHandler is an event handler for Hello events.
type helloEventTypeHandler func(*Client, *EventHello)

// Type returns the event type for Hello events.
func (eventTypeHandler helloEventTypeHandler) Type() string {
	return helloEventType
}

// New returns a new instance of Hello.
func (eventTypeHandler helloEventTypeHandler) New() interface{} {
	return &EventHello{}
}

// Handle is the handler for Hello events.
func (eventTypeHandler helloEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventHello)
	err := c.onHello(e)
	return err
}

// inviteCreateEventTypeHandler is an event handler for InviteCreate events.
type inviteCreateEventTypeHandler func(*Client, *EventInviteCreate)

// Type returns the event type for InviteCreate events.
func (eventTypeHandler inviteCreateEventTypeHandler) Type() string {
	return inviteCreateEventType
}

// New returns a new instance of InviteCreate.
func (eventTypeHandler inviteCreateEventTypeHandler) New() interface{} {
	return &EventInviteCreate{}
}

// Handle is the handler for InviteCreate events.
func (eventTypeHandler inviteCreateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventInviteCreate)
	err := c.onInviteCreate(e)
	return err
}

// inviteDeleteEventTypeHandler is an event handler for InviteDelete events.
type inviteDeleteEventTypeHandler func(*Client, *EventInviteDelete)

// Type returns the event type for InviteDelete events.
func (eventTypeHandler inviteDeleteEventTypeHandler) Type() string {
	return inviteDeleteEventType
}

// New returns a new instance of InviteDelete.
func (eventTypeHandler inviteDeleteEventTypeHandler) New() interface{} {
	return &EventInviteDelete{}
}

// Handle is the handler for InviteDelete events.
func (eventTypeHandler inviteDeleteEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventInviteDelete)
	err := c.onInviteDelete(e)
	return err
}

// messageCreateEventTypeHandler is an event handler for MessageCreate events.
type messageCreateEventTypeHandler func(*Client, *EventMessageCreate)

// Type returns the event type for MessageCreate events.
func (eventTypeHandler messageCreateEventTypeHandler) Type() string {
	return messageCreateEventType
}

// New returns a new instance of MessageCreate.
func (eventTypeHandler messageCreateEventTypeHandler) New() interface{} {
	return &EventMessageCreate{}
}

// Handle is the handler for MessageCreate events.
func (eventTypeHandler messageCreateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageCreate)
	err := c.onMessageCreate(e)
	return err
}

// messageDeleteEventTypeHandler is an event handler for MessageDelete events.
type messageDeleteEventTypeHandler func(*Client, *EventMessageDelete)

// Type returns the event type for MessageDelete events.
func (eventTypeHandler messageDeleteEventTypeHandler) Type() string {
	return messageDeleteEventType
}

// New returns a new instance of MessageDelete.
func (eventTypeHandler messageDeleteEventTypeHandler) New() interface{} {
	return &EventMessageDelete{}
}

// Handle is the handler for MessageDelete events.
func (eventTypeHandler messageDeleteEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageDelete)
	err := c.onMessageDelete(e)
	return err
}

// messageDeleteBulkEventTypeHandler is an event handler for MessageDeleteBulk events.
type messageDeleteBulkEventTypeHandler func(*Client, *EventMessageDeleteBulk)

// Type returns the event type for MessageDeleteBulk events.
func (eventTypeHandler messageDeleteBulkEventTypeHandler) Type() string {
	return messageDeleteBulkEventType
}

// New returns a new instance of MessageDeleteBulk.
func (eventTypeHandler messageDeleteBulkEventTypeHandler) New() interface{} {
	return &EventMessageDeleteBulk{}
}

// Handle is the handler for MessageDeleteBulk events.
func (eventTypeHandler messageDeleteBulkEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageDeleteBulk)
	err := c.onMessageDeleteBulk(e)
	return err
}

// messageReactionAddEventTypeHandler is an event handler for MessageReactionAdd events.
type messageReactionAddEventTypeHandler func(*Client, *EventMessageReactionAdd)

// Type returns the event type for MessageReactionAdd events.
func (eventTypeHandler messageReactionAddEventTypeHandler) Type() string {
	return messageReactionAddEventType
}

// New returns a new instance of MessageReactionAdd.
func (eventTypeHandler messageReactionAddEventTypeHandler) New() interface{} {
	return &EventMessageReactionAdd{}
}

// Handle is the handler for MessageReactionAdd events.
func (eventTypeHandler messageReactionAddEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageReactionAdd)
	err := c.onMessageReactionAdd(e)
	return err
}

// messageReactionRemoveEventTypeHandler is an event handler for MessageReactionRemove events.
type messageReactionRemoveEventTypeHandler func(*Client, *EventMessageReactionRemove)

// Type returns the event type for MessageReactionRemove events.
func (eventTypeHandler messageReactionRemoveEventTypeHandler) Type() string {
	return messageReactionRemoveEventType
}

// New returns a new instance of MessageReactionRemove.
func (eventTypeHandler messageReactionRemoveEventTypeHandler) New() interface{} {
	return &EventMessageReactionRemove{}
}

// Handle is the handler for MessageReactionRemove events.
func (eventTypeHandler messageReactionRemoveEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageReactionRemove)
	err := c.onMessageReactionRemove(e)
	return err
}

// messageReactionRemoveAllEventTypeHandler is an event handler for MessageReactionRemoveAll events.
type messageReactionRemoveAllEventTypeHandler func(*Client, *EventMessageReactionRemoveAll)

// Type returns the event type for MessageReactionRemoveAll events.
func (eventTypeHandler messageReactionRemoveAllEventTypeHandler) Type() string {
	return messageReactionRemoveAllEventType
}

// New returns a new instance of MessageReactionRemoveAll.
func (eventTypeHandler messageReactionRemoveAllEventTypeHandler) New() interface{} {
	return &EventMessageReactionRemoveAll{}
}

// Handle is the handler for MessageReactionRemoveAll events.
func (eventTypeHandler messageReactionRemoveAllEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageReactionRemoveAll)
	err := c.onMessageReactionRemoveAll(e)
	return err
}

// messageReactionRemoveEmojiEventTypeHandler is an event handler for MessageReactionRemoveEmoji events.
type messageReactionRemoveEmojiEventTypeHandler func(*Client, *EventMessageReactionRemoveEmoji)

// Type returns the event type for MessageReactionRemoveEmoji events.
func (eventTypeHandler messageReactionRemoveEmojiEventTypeHandler) Type() string {
	return messageReactionRemoveEmojiEventType
}

// New returns a new instance of MessageReactionRemoveEmoji.
func (eventTypeHandler messageReactionRemoveEmojiEventTypeHandler) New() interface{} {
	return &EventMessageReactionRemoveEmoji{}
}

// Handle is the handler for MessageReactionRemoveEmoji events.
func (eventTypeHandler messageReactionRemoveEmojiEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageReactionRemoveEmoji)
	err := c.onMessageReactionRemoveEmoji(e)
	return err
}

// messageUpdateEventTypeHandler is an event handler for MessageUpdate events.
type messageUpdateEventTypeHandler func(*Client, *EventMessageUpdate)

// Type returns the event type for MessageUpdate events.
func (eventTypeHandler messageUpdateEventTypeHandler) Type() string {
	return messageUpdateEventType
}

// New returns a new instance of MessageUpdate.
func (eventTypeHandler messageUpdateEventTypeHandler) New() interface{} {
	return &EventMessageUpdate{}
}

// Handle is the handler for MessageUpdate events.
func (eventTypeHandler messageUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventMessageUpdate)
	err := c.onMessageUpdate(e)
	return err
}

// presenceUpdateEventTypeHandler is an event handler for PresenceUpdate events.
type presenceUpdateEventTypeHandler func(*Client, *EventPresenceUpdate)

// Type returns the event type for PresenceUpdate events.
func (eventTypeHandler presenceUpdateEventTypeHandler) Type() string {
	return presenceUpdateEventType
}

// New returns a new instance of PresenceUpdate.
func (eventTypeHandler presenceUpdateEventTypeHandler) New() interface{} {
	return &EventPresenceUpdate{}
}

// Handle is the handler for PresenceUpdate events.
func (eventTypeHandler presenceUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventPresenceUpdate)
	err := c.onPresenceUpdate(e)
	return err
}

// readyEventTypeHandler is an event handler for Ready events.
type readyEventTypeHandler func(*Client, *EventReady)

// Type returns the event type for Ready events.
func (eventTypeHandler readyEventTypeHandler) Type() string {
	return readyEventType
}

// New returns a new instance of Ready.
func (eventTypeHandler readyEventTypeHandler) New() interface{} {
	return &EventReady{}
}

// Handle is the handler for Ready events.
func (eventTypeHandler readyEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventReady)
	err := c.onReady(e)
	return err
}

// reconnectEventTypeHandler is an event handler for Reconnect events.
type reconnectEventTypeHandler func(*Client, *EventReconnect)

// Type returns the event type for Reconnect events.
func (eventTypeHandler reconnectEventTypeHandler) Type() string {
	return reconnectEventType
}

// New returns a new instance of Reconnect.
func (eventTypeHandler reconnectEventTypeHandler) New() interface{} {
	return &EventReconnect{}
}

// Handle is the handler for Reconnect events.
func (eventTypeHandler reconnectEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventReconnect)
	err := c.onReconnect(e)
	return err
}

// resumedEventTypeHandler is an event handler for Resumed events.
type resumedEventTypeHandler func(*Client, *EventResumed)

// Type returns the event type for Resumed events.
func (eventTypeHandler resumedEventTypeHandler) Type() string {
	return resumedEventType
}

// New returns a new instance of Resumed.
func (eventTypeHandler resumedEventTypeHandler) New() interface{} {
	return &EventResumed{}
}

// Handle is the handler for Resumed events.
func (eventTypeHandler resumedEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventResumed)
	err := c.onResumed(e)
	return err
}

// typingStartEventTypeHandler is an event handler for TypingStart events.
type typingStartEventTypeHandler func(*Client, *EventTypingStart)

// Type returns the event type for TypingStart events.
func (eventTypeHandler typingStartEventTypeHandler) Type() string {
	return typingStartEventType
}

// New returns a new instance of TypingStart.
func (eventTypeHandler typingStartEventTypeHandler) New() interface{} {
	return &EventTypingStart{}
}

// Handle is the handler for TypingStart events.
func (eventTypeHandler typingStartEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventTypingStart)
	err := c.onTypingStart(e)
	return err
}

// userUpdateEventTypeHandler is an event handler for UserUpdate events.
type userUpdateEventTypeHandler func(*Client, *EventUserUpdate)

// Type returns the event type for UserUpdate events.
func (eventTypeHandler userUpdateEventTypeHandler) Type() string {
	return userUpdateEventType
}

// New returns a new instance of UserUpdate.
func (eventTypeHandler userUpdateEventTypeHandler) New() interface{} {
	return &EventUserUpdate{}
}

// Handle is the handler for UserUpdate events.
func (eventTypeHandler userUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventUserUpdate)
	err := c.onUserUpdate(e)
	return err
}

// voiceServerUpdateEventTypeHandler is an event handler for VoiceServerUpdate events.
type voiceServerUpdateEventTypeHandler func(*Client, *EventVoiceServerUpdate)

// Type returns the event type for VoiceServerUpdate events.
func (eventTypeHandler voiceServerUpdateEventTypeHandler) Type() string {
	return voiceServerUpdateEventType
}

// New returns a new instance of VoiceServerUpdate.
func (eventTypeHandler voiceServerUpdateEventTypeHandler) New() interface{} {
	return &EventVoiceServerUpdate{}
}

// Handle is the handler for VoiceServerUpdate events.
func (eventTypeHandler voiceServerUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventVoiceServerUpdate)
	err := c.onVoiceServerUpdate(e)
	return err
}

// voiceStateUpdateEventTypeHandler is an event handler for VoiceStateUpdate events.
type voiceStateUpdateEventTypeHandler func(*Client, *EventVoiceStateUpdate)

// Type returns the event type for VoiceStateUpdate events.
func (eventTypeHandler voiceStateUpdateEventTypeHandler) Type() string {
	return voiceStateUpdateEventType
}

// New returns a new instance of VoiceStateUpdate.
func (eventTypeHandler voiceStateUpdateEventTypeHandler) New() interface{} {
	return &EventVoiceStateUpdate{}
}

// Handle is the handler for VoiceStateUpdate events.
func (eventTypeHandler voiceStateUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventVoiceStateUpdate)
	err := c.onVoiceStateUpdate(e)
	return err
}

// webhooksUpdateEventTypeHandler is an event handler for WebhooksUpdate events.
type webhooksUpdateEventTypeHandler func(*Client, *EventWebhooksUpdate)

// Type returns the event type for WebhooksUpdate events.
func (eventTypeHandler webhooksUpdateEventTypeHandler) Type() string {
	return webhooksUpdateEventType
}

// New returns a new instance of WebhooksUpdate.
func (eventTypeHandler webhooksUpdateEventTypeHandler) New() interface{} {
	return &EventWebhooksUpdate{}
}

// Handle is the handler for WebhooksUpdate events.
func (eventTypeHandler webhooksUpdateEventTypeHandler) Handle(c *Client, i interface{}) error {
	e := i.(*EventWebhooksUpdate)
	err := c.onWebhooksUpdate(e)
	return err
}

func handlerForInterface(handler interface{}) EventTypeHandler {
	switch v := handler.(type) {
	case func(*Client, *EventChannelCreate):
		return channelCreateEventTypeHandler(v)
	case func(*Client, *EventChannelDelete):
		return channelDeleteEventTypeHandler(v)
	case func(*Client, *EventChannelPinsUpdate):
		return channelPinsUpdateEventTypeHandler(v)
	case func(*Client, *EventChannelUpdate):
		return channelUpdateEventTypeHandler(v)
	case func(*Client, *EventGuildBanAdd):
		return guildBanAddEventTypeHandler(v)
	case func(*Client, *EventGuildBanRemove):
		return guildBanRemoveEventTypeHandler(v)
	case func(*Client, *EventGuildCreate):
		return guildCreateEventTypeHandler(v)
	case func(*Client, *EventGuildDelete):
		return guildDeleteEventTypeHandler(v)
	case func(*Client, *EventGuildEmojisUpdate):
		return guildEmojisUpdateEventTypeHandler(v)
	case func(*Client, *EventGuildIntegrationsUpdate):
		return guildIntegrationsUpdateEventTypeHandler(v)
	case func(*Client, *EventGuildMemberAdd):
		return guildMemberAddEventTypeHandler(v)
	case func(*Client, *EventGuildMemberRemove):
		return guildMemberRemoveEventTypeHandler(v)
	case func(*Client, *EventGuildMemberUpdate):
		return guildMemberUpdateEventTypeHandler(v)
	case func(*Client, *EventGuildRoleCreate):
		return guildRoleCreateEventTypeHandler(v)
	case func(*Client, *EventGuildRoleDelete):
		return guildRoleDeleteEventTypeHandler(v)
	case func(*Client, *EventGuildRoleUpdate):
		return guildRoleUpdateEventTypeHandler(v)
	case func(*Client, *EventGuildUpdate):
		return guildUpdateEventTypeHandler(v)
	case func(*Client, *EventHello):
		return helloEventTypeHandler(v)
	case func(*Client, *EventInviteCreate):
		return inviteCreateEventTypeHandler(v)
	case func(*Client, *EventInviteDelete):
		return inviteDeleteEventTypeHandler(v)
	case func(*Client, *EventMessageCreate):
		return messageCreateEventTypeHandler(v)
	case func(*Client, *EventMessageDelete):
		return messageDeleteEventTypeHandler(v)
	case func(*Client, *EventMessageDeleteBulk):
		return messageDeleteBulkEventTypeHandler(v)
	case func(*Client, *EventMessageReactionAdd):
		return messageReactionAddEventTypeHandler(v)
	case func(*Client, *EventMessageReactionRemove):
		return messageReactionRemoveEventTypeHandler(v)
	case func(*Client, *EventMessageReactionRemoveAll):
		return messageReactionRemoveAllEventTypeHandler(v)
	case func(*Client, *EventMessageReactionRemoveEmoji):
		return messageReactionRemoveEmojiEventTypeHandler(v)
	case func(*Client, *EventMessageUpdate):
		return messageUpdateEventTypeHandler(v)
	case func(*Client, *EventPresenceUpdate):
		return presenceUpdateEventTypeHandler(v)
	case func(*Client, *EventReady):
		return readyEventTypeHandler(v)
	case func(*Client, *EventReconnect):
		return reconnectEventTypeHandler(v)
	case func(*Client, *EventResumed):
		return resumedEventTypeHandler(v)
	case func(*Client, *EventTypingStart):
		return typingStartEventTypeHandler(v)
	case func(*Client, *EventUserUpdate):
		return userUpdateEventTypeHandler(v)
	case func(*Client, *EventVoiceServerUpdate):
		return voiceServerUpdateEventTypeHandler(v)
	case func(*Client, *EventVoiceStateUpdate):
		return voiceStateUpdateEventTypeHandler(v)
	case func(*Client, *EventWebhooksUpdate):
		return webhooksUpdateEventTypeHandler(v)
	}
	return nil
}

// EventTypeHandler represents any EventTypeHandler
type EventTypeHandler interface {
	Type() string
	New() interface{}
	Handle(c *Client, i interface{}) error
}

var eventTypeHandlers = map[string]EventTypeHandler{}

func addEventTypeHandler(eventTypeHandler EventTypeHandler) {
	eventTypeHandlers[eventTypeHandler.Type()] = eventTypeHandler
}

// AddEventTypeHandlers maps all event handlers
func AddEventTypeHandlers() {
	addEventTypeHandler(channelCreateEventTypeHandler(nil))
	addEventTypeHandler(channelDeleteEventTypeHandler(nil))
	addEventTypeHandler(channelPinsUpdateEventTypeHandler(nil))
	addEventTypeHandler(channelUpdateEventTypeHandler(nil))
	addEventTypeHandler(guildBanAddEventTypeHandler(nil))
	addEventTypeHandler(guildBanRemoveEventTypeHandler(nil))
	addEventTypeHandler(guildCreateEventTypeHandler(nil))
	addEventTypeHandler(guildDeleteEventTypeHandler(nil))
	addEventTypeHandler(guildEmojisUpdateEventTypeHandler(nil))
	addEventTypeHandler(guildIntegrationsUpdateEventTypeHandler(nil))
	addEventTypeHandler(guildMemberAddEventTypeHandler(nil))
	addEventTypeHandler(guildMemberRemoveEventTypeHandler(nil))
	addEventTypeHandler(guildMemberUpdateEventTypeHandler(nil))
	addEventTypeHandler(guildRoleCreateEventTypeHandler(nil))
	addEventTypeHandler(guildRoleDeleteEventTypeHandler(nil))
	addEventTypeHandler(guildRoleUpdateEventTypeHandler(nil))
	addEventTypeHandler(guildUpdateEventTypeHandler(nil))
	addEventTypeHandler(helloEventTypeHandler(nil))
	addEventTypeHandler(inviteCreateEventTypeHandler(nil))
	addEventTypeHandler(inviteDeleteEventTypeHandler(nil))
	addEventTypeHandler(messageCreateEventTypeHandler(nil))
	addEventTypeHandler(messageDeleteEventTypeHandler(nil))
	addEventTypeHandler(messageDeleteBulkEventTypeHandler(nil))
	addEventTypeHandler(messageReactionAddEventTypeHandler(nil))
	addEventTypeHandler(messageReactionRemoveEventTypeHandler(nil))
	addEventTypeHandler(messageReactionRemoveAllEventTypeHandler(nil))
	addEventTypeHandler(messageReactionRemoveEmojiEventTypeHandler(nil))
	addEventTypeHandler(messageUpdateEventTypeHandler(nil))
	addEventTypeHandler(presenceUpdateEventTypeHandler(nil))
	addEventTypeHandler(readyEventTypeHandler(nil))
	addEventTypeHandler(reconnectEventTypeHandler(nil))
	addEventTypeHandler(resumedEventTypeHandler(nil))
	addEventTypeHandler(typingStartEventTypeHandler(nil))
	addEventTypeHandler(userUpdateEventTypeHandler(nil))
	addEventTypeHandler(voiceServerUpdateEventTypeHandler(nil))
	addEventTypeHandler(voiceStateUpdateEventTypeHandler(nil))
	addEventTypeHandler(webhooksUpdateEventTypeHandler(nil))
}
