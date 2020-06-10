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

// channelCreateEventHandler is an event handler for ChannelCreate events.
type channelCreateEventHandler func(*Client, *EventChannelCreate)

// Type returns the event type for ChannelCreate events.
func (eventhandler channelCreateEventHandler) Type() string {
	return channelCreateEventType
}

// New returns a new instance of ChannelCreate.
func (eventhandler channelCreateEventHandler) New() interface{} {
	return &EventChannelCreate{}
}

// Handle is the handler for ChannelCreate events.
func (eventhandler channelCreateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelCreate); ok {
		eventhandler(c, t)
	}
}

// channelDeleteEventHandler is an event handler for ChannelDelete events.
type channelDeleteEventHandler func(*Client, *EventChannelDelete)

// Type returns the event type for ChannelDelete events.
func (eventhandler channelDeleteEventHandler) Type() string {
	return channelDeleteEventType
}

// New returns a new instance of ChannelDelete.
func (eventhandler channelDeleteEventHandler) New() interface{} {
	return &EventChannelDelete{}
}

// Handle is the handler for ChannelDelete events.
func (eventhandler channelDeleteEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelDelete); ok {
		eventhandler(c, t)
	}
}

// channelPinsUpdateEventHandler is an event handler for ChannelPinsUpdate events.
type channelPinsUpdateEventHandler func(*Client, *EventChannelPinsUpdate)

// Type returns the event type for ChannelPinsUpdate events.
func (eventhandler channelPinsUpdateEventHandler) Type() string {
	return channelPinsUpdateEventType
}

// New returns a new instance of ChannelPinsUpdate.
func (eventhandler channelPinsUpdateEventHandler) New() interface{} {
	return &EventChannelPinsUpdate{}
}

// Handle is the handler for ChannelPinsUpdate events.
func (eventhandler channelPinsUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelPinsUpdate); ok {
		eventhandler(c, t)
	}
}

// channelUpdateEventHandler is an event handler for ChannelUpdate events.
type channelUpdateEventHandler func(*Client, *EventChannelUpdate)

// Type returns the event type for ChannelUpdate events.
func (eventhandler channelUpdateEventHandler) Type() string {
	return channelUpdateEventType
}

// New returns a new instance of ChannelUpdate.
func (eventhandler channelUpdateEventHandler) New() interface{} {
	return &EventChannelUpdate{}
}

// Handle is the handler for ChannelUpdate events.
func (eventhandler channelUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelUpdate); ok {
		eventhandler(c, t)
	}
}

// guildBanAddEventHandler is an event handler for GuildBanAdd events.
type guildBanAddEventHandler func(*Client, *EventGuildBanAdd)

// Type returns the event type for GuildBanAdd events.
func (eventhandler guildBanAddEventHandler) Type() string {
	return guildBanAddEventType
}

// New returns a new instance of GuildBanAdd.
func (eventhandler guildBanAddEventHandler) New() interface{} {
	return &EventGuildBanAdd{}
}

// Handle is the handler for GuildBanAdd events.
func (eventhandler guildBanAddEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildBanAdd); ok {
		eventhandler(c, t)
	}
}

// guildBanRemoveEventHandler is an event handler for GuildBanRemove events.
type guildBanRemoveEventHandler func(*Client, *EventGuildBanRemove)

// Type returns the event type for GuildBanRemove events.
func (eventhandler guildBanRemoveEventHandler) Type() string {
	return guildBanRemoveEventType
}

// New returns a new instance of GuildBanRemove.
func (eventhandler guildBanRemoveEventHandler) New() interface{} {
	return &EventGuildBanRemove{}
}

// Handle is the handler for GuildBanRemove events.
func (eventhandler guildBanRemoveEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildBanRemove); ok {
		eventhandler(c, t)
	}
}

// guildCreateEventHandler is an event handler for GuildCreate events.
type guildCreateEventHandler func(*Client, *EventGuildCreate)

// Type returns the event type for GuildCreate events.
func (eventhandler guildCreateEventHandler) Type() string {
	return guildCreateEventType
}

// New returns a new instance of GuildCreate.
func (eventhandler guildCreateEventHandler) New() interface{} {
	return &EventGuildCreate{}
}

// Handle is the handler for GuildCreate events.
func (eventhandler guildCreateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildCreate); ok {
		eventhandler(c, t)
	}
}

// guildDeleteEventHandler is an event handler for GuildDelete events.
type guildDeleteEventHandler func(*Client, *EventGuildDelete)

// Type returns the event type for GuildDelete events.
func (eventhandler guildDeleteEventHandler) Type() string {
	return guildDeleteEventType
}

// New returns a new instance of GuildDelete.
func (eventhandler guildDeleteEventHandler) New() interface{} {
	return &EventGuildDelete{}
}

// Handle is the handler for GuildDelete events.
func (eventhandler guildDeleteEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildDelete); ok {
		eventhandler(c, t)
	}
}

// guildEmojisUpdateEventHandler is an event handler for GuildEmojisUpdate events.
type guildEmojisUpdateEventHandler func(*Client, *EventGuildEmojisUpdate)

// Type returns the event type for GuildEmojisUpdate events.
func (eventhandler guildEmojisUpdateEventHandler) Type() string {
	return guildEmojisUpdateEventType
}

// New returns a new instance of GuildEmojisUpdate.
func (eventhandler guildEmojisUpdateEventHandler) New() interface{} {
	return &EventGuildEmojisUpdate{}
}

// Handle is the handler for GuildEmojisUpdate events.
func (eventhandler guildEmojisUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildEmojisUpdate); ok {
		eventhandler(c, t)
	}
}

// guildIntegrationsUpdateEventHandler is an event handler for GuildIntegrationsUpdate events.
type guildIntegrationsUpdateEventHandler func(*Client, *EventGuildIntegrationsUpdate)

// Type returns the event type for GuildIntegrationsUpdate events.
func (eventhandler guildIntegrationsUpdateEventHandler) Type() string {
	return guildIntegrationsUpdateEventType
}

// New returns a new instance of GuildIntegrationsUpdate.
func (eventhandler guildIntegrationsUpdateEventHandler) New() interface{} {
	return &EventGuildIntegrationsUpdate{}
}

// Handle is the handler for GuildIntegrationsUpdate events.
func (eventhandler guildIntegrationsUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildIntegrationsUpdate); ok {
		eventhandler(c, t)
	}
}

// guildMemberAddEventHandler is an event handler for GuildMemberAdd events.
type guildMemberAddEventHandler func(*Client, *EventGuildMemberAdd)

// Type returns the event type for GuildMemberAdd events.
func (eventhandler guildMemberAddEventHandler) Type() string {
	return guildMemberAddEventType
}

// New returns a new instance of GuildMemberAdd.
func (eventhandler guildMemberAddEventHandler) New() interface{} {
	return &EventGuildMemberAdd{}
}

// Handle is the handler for GuildMemberAdd events.
func (eventhandler guildMemberAddEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildMemberAdd); ok {
		eventhandler(c, t)
	}
}

// guildMemberRemoveEventHandler is an event handler for GuildMemberRemove events.
type guildMemberRemoveEventHandler func(*Client, *EventGuildMemberRemove)

// Type returns the event type for GuildMemberRemove events.
func (eventhandler guildMemberRemoveEventHandler) Type() string {
	return guildMemberRemoveEventType
}

// New returns a new instance of GuildMemberRemove.
func (eventhandler guildMemberRemoveEventHandler) New() interface{} {
	return &EventGuildMemberRemove{}
}

// Handle is the handler for GuildMemberRemove events.
func (eventhandler guildMemberRemoveEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildMemberRemove); ok {
		eventhandler(c, t)
	}
}

// guildMemberUpdateEventHandler is an event handler for GuildMemberUpdate events.
type guildMemberUpdateEventHandler func(*Client, *EventGuildMemberUpdate)

// Type returns the event type for GuildMemberUpdate events.
func (eventhandler guildMemberUpdateEventHandler) Type() string {
	return guildMemberUpdateEventType
}

// New returns a new instance of GuildMemberUpdate.
func (eventhandler guildMemberUpdateEventHandler) New() interface{} {
	return &EventGuildMemberUpdate{}
}

// Handle is the handler for GuildMemberUpdate events.
func (eventhandler guildMemberUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildMemberUpdate); ok {
		eventhandler(c, t)
	}
}

// guildRoleCreateEventHandler is an event handler for GuildRoleCreate events.
type guildRoleCreateEventHandler func(*Client, *EventGuildRoleCreate)

// Type returns the event type for GuildRoleCreate events.
func (eventhandler guildRoleCreateEventHandler) Type() string {
	return guildRoleCreateEventType
}

// New returns a new instance of GuildRoleCreate.
func (eventhandler guildRoleCreateEventHandler) New() interface{} {
	return &EventGuildRoleCreate{}
}

// Handle is the handler for GuildRoleCreate events.
func (eventhandler guildRoleCreateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildRoleCreate); ok {
		eventhandler(c, t)
	}
}

// guildRoleDeleteEventHandler is an event handler for GuildRoleDelete events.
type guildRoleDeleteEventHandler func(*Client, *EventGuildRoleDelete)

// Type returns the event type for GuildRoleDelete events.
func (eventhandler guildRoleDeleteEventHandler) Type() string {
	return guildRoleDeleteEventType
}

// New returns a new instance of GuildRoleDelete.
func (eventhandler guildRoleDeleteEventHandler) New() interface{} {
	return &EventGuildRoleDelete{}
}

// Handle is the handler for GuildRoleDelete events.
func (eventhandler guildRoleDeleteEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildRoleDelete); ok {
		eventhandler(c, t)
	}
}

// guildRoleUpdateEventHandler is an event handler for GuildRoleUpdate events.
type guildRoleUpdateEventHandler func(*Client, *EventGuildRoleUpdate)

// Type returns the event type for GuildRoleUpdate events.
func (eventhandler guildRoleUpdateEventHandler) Type() string {
	return guildRoleUpdateEventType
}

// New returns a new instance of GuildRoleUpdate.
func (eventhandler guildRoleUpdateEventHandler) New() interface{} {
	return &EventGuildRoleUpdate{}
}

// Handle is the handler for GuildRoleUpdate events.
func (eventhandler guildRoleUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildRoleUpdate); ok {
		eventhandler(c, t)
	}
}

// guildUpdateEventHandler is an event handler for GuildUpdate events.
type guildUpdateEventHandler func(*Client, *EventGuildUpdate)

// Type returns the event type for GuildUpdate events.
func (eventhandler guildUpdateEventHandler) Type() string {
	return guildUpdateEventType
}

// New returns a new instance of GuildUpdate.
func (eventhandler guildUpdateEventHandler) New() interface{} {
	return &EventGuildUpdate{}
}

// Handle is the handler for GuildUpdate events.
func (eventhandler guildUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildUpdate); ok {
		eventhandler(c, t)
	}
}

// helloEventHandler is an event handler for Hello events.
type helloEventHandler func(*Client, *EventHello)

// Type returns the event type for Hello events.
func (eventhandler helloEventHandler) Type() string {
	return helloEventType
}

// New returns a new instance of Hello.
func (eventhandler helloEventHandler) New() interface{} {
	return &EventHello{}
}

// Handle is the handler for Hello events.
func (eventhandler helloEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventHello); ok {
		eventhandler(c, t)
	}
}

// inviteCreateEventHandler is an event handler for InviteCreate events.
type inviteCreateEventHandler func(*Client, *EventInviteCreate)

// Type returns the event type for InviteCreate events.
func (eventhandler inviteCreateEventHandler) Type() string {
	return inviteCreateEventType
}

// New returns a new instance of InviteCreate.
func (eventhandler inviteCreateEventHandler) New() interface{} {
	return &EventInviteCreate{}
}

// Handle is the handler for InviteCreate events.
func (eventhandler inviteCreateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventInviteCreate); ok {
		eventhandler(c, t)
	}
}

// inviteDeleteEventHandler is an event handler for InviteDelete events.
type inviteDeleteEventHandler func(*Client, *EventInviteDelete)

// Type returns the event type for InviteDelete events.
func (eventhandler inviteDeleteEventHandler) Type() string {
	return inviteDeleteEventType
}

// New returns a new instance of InviteDelete.
func (eventhandler inviteDeleteEventHandler) New() interface{} {
	return &EventInviteDelete{}
}

// Handle is the handler for InviteDelete events.
func (eventhandler inviteDeleteEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventInviteDelete); ok {
		eventhandler(c, t)
	}
}

// messageCreateEventHandler is an event handler for MessageCreate events.
type messageCreateEventHandler func(*Client, *EventMessageCreate)

// Type returns the event type for MessageCreate events.
func (eventhandler messageCreateEventHandler) Type() string {
	return messageCreateEventType
}

// New returns a new instance of MessageCreate.
func (eventhandler messageCreateEventHandler) New() interface{} {
	return &EventMessageCreate{}
}

// Handle is the handler for MessageCreate events.
func (eventhandler messageCreateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageCreate); ok {
		eventhandler(c, t)
	}
}

// messageDeleteEventHandler is an event handler for MessageDelete events.
type messageDeleteEventHandler func(*Client, *EventMessageDelete)

// Type returns the event type for MessageDelete events.
func (eventhandler messageDeleteEventHandler) Type() string {
	return messageDeleteEventType
}

// New returns a new instance of MessageDelete.
func (eventhandler messageDeleteEventHandler) New() interface{} {
	return &EventMessageDelete{}
}

// Handle is the handler for MessageDelete events.
func (eventhandler messageDeleteEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageDelete); ok {
		eventhandler(c, t)
	}
}

// messageDeleteBulkEventHandler is an event handler for MessageDeleteBulk events.
type messageDeleteBulkEventHandler func(*Client, *EventMessageDeleteBulk)

// Type returns the event type for MessageDeleteBulk events.
func (eventhandler messageDeleteBulkEventHandler) Type() string {
	return messageDeleteBulkEventType
}

// New returns a new instance of MessageDeleteBulk.
func (eventhandler messageDeleteBulkEventHandler) New() interface{} {
	return &EventMessageDeleteBulk{}
}

// Handle is the handler for MessageDeleteBulk events.
func (eventhandler messageDeleteBulkEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageDeleteBulk); ok {
		eventhandler(c, t)
	}
}

// messageReactionAddEventHandler is an event handler for MessageReactionAdd events.
type messageReactionAddEventHandler func(*Client, *EventMessageReactionAdd)

// Type returns the event type for MessageReactionAdd events.
func (eventhandler messageReactionAddEventHandler) Type() string {
	return messageReactionAddEventType
}

// New returns a new instance of MessageReactionAdd.
func (eventhandler messageReactionAddEventHandler) New() interface{} {
	return &EventMessageReactionAdd{}
}

// Handle is the handler for MessageReactionAdd events.
func (eventhandler messageReactionAddEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageReactionAdd); ok {
		eventhandler(c, t)
	}
}

// messageReactionRemoveEventHandler is an event handler for MessageReactionRemove events.
type messageReactionRemoveEventHandler func(*Client, *EventMessageReactionRemove)

// Type returns the event type for MessageReactionRemove events.
func (eventhandler messageReactionRemoveEventHandler) Type() string {
	return messageReactionRemoveEventType
}

// New returns a new instance of MessageReactionRemove.
func (eventhandler messageReactionRemoveEventHandler) New() interface{} {
	return &EventMessageReactionRemove{}
}

// Handle is the handler for MessageReactionRemove events.
func (eventhandler messageReactionRemoveEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageReactionRemove); ok {
		eventhandler(c, t)
	}
}

// messageReactionRemoveAllEventHandler is an event handler for MessageReactionRemoveAll events.
type messageReactionRemoveAllEventHandler func(*Client, *EventMessageReactionRemoveAll)

// Type returns the event type for MessageReactionRemoveAll events.
func (eventhandler messageReactionRemoveAllEventHandler) Type() string {
	return messageReactionRemoveAllEventType
}

// New returns a new instance of MessageReactionRemoveAll.
func (eventhandler messageReactionRemoveAllEventHandler) New() interface{} {
	return &EventMessageReactionRemoveAll{}
}

// Handle is the handler for MessageReactionRemoveAll events.
func (eventhandler messageReactionRemoveAllEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageReactionRemoveAll); ok {
		eventhandler(c, t)
	}
}

// messageReactionRemoveEmojiEventHandler is an event handler for MessageReactionRemoveEmoji events.
type messageReactionRemoveEmojiEventHandler func(*Client, *EventMessageReactionRemoveEmoji)

// Type returns the event type for MessageReactionRemoveEmoji events.
func (eventhandler messageReactionRemoveEmojiEventHandler) Type() string {
	return messageReactionRemoveEmojiEventType
}

// New returns a new instance of MessageReactionRemoveEmoji.
func (eventhandler messageReactionRemoveEmojiEventHandler) New() interface{} {
	return &EventMessageReactionRemoveEmoji{}
}

// Handle is the handler for MessageReactionRemoveEmoji events.
func (eventhandler messageReactionRemoveEmojiEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageReactionRemoveEmoji); ok {
		eventhandler(c, t)
	}
}

// messageUpdateEventHandler is an event handler for MessageUpdate events.
type messageUpdateEventHandler func(*Client, *EventMessageUpdate)

// Type returns the event type for MessageUpdate events.
func (eventhandler messageUpdateEventHandler) Type() string {
	return messageUpdateEventType
}

// New returns a new instance of MessageUpdate.
func (eventhandler messageUpdateEventHandler) New() interface{} {
	return &EventMessageUpdate{}
}

// Handle is the handler for MessageUpdate events.
func (eventhandler messageUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageUpdate); ok {
		eventhandler(c, t)
	}
}

// presenceUpdateEventHandler is an event handler for PresenceUpdate events.
type presenceUpdateEventHandler func(*Client, *EventPresenceUpdate)

// Type returns the event type for PresenceUpdate events.
func (eventhandler presenceUpdateEventHandler) Type() string {
	return presenceUpdateEventType
}

// New returns a new instance of PresenceUpdate.
func (eventhandler presenceUpdateEventHandler) New() interface{} {
	return &EventPresenceUpdate{}
}

// Handle is the handler for PresenceUpdate events.
func (eventhandler presenceUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventPresenceUpdate); ok {
		eventhandler(c, t)
	}
}

// readyEventHandler is an event handler for Ready events.
type readyEventHandler func(*Client, *EventReady)

// Type returns the event type for Ready events.
func (eventhandler readyEventHandler) Type() string {
	return readyEventType
}

// New returns a new instance of Ready.
func (eventhandler readyEventHandler) New() interface{} {
	return &EventReady{}
}

// Handle is the handler for Ready events.
func (eventhandler readyEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventReady); ok {
		eventhandler(c, t)
	}
}

// reconnectEventHandler is an event handler for Reconnect events.
type reconnectEventHandler func(*Client, *EventReconnect)

// Type returns the event type for Reconnect events.
func (eventhandler reconnectEventHandler) Type() string {
	return reconnectEventType
}

// New returns a new instance of Reconnect.
func (eventhandler reconnectEventHandler) New() interface{} {
	return &EventReconnect{}
}

// Handle is the handler for Reconnect events.
func (eventhandler reconnectEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventReconnect); ok {
		eventhandler(c, t)
	}
}

// resumedEventHandler is an event handler for Resumed events.
type resumedEventHandler func(*Client, *EventResumed)

// Type returns the event type for Resumed events.
func (eventhandler resumedEventHandler) Type() string {
	return resumedEventType
}

// New returns a new instance of Resumed.
func (eventhandler resumedEventHandler) New() interface{} {
	return &EventResumed{}
}

// Handle is the handler for Resumed events.
func (eventhandler resumedEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventResumed); ok {
		eventhandler(c, t)
	}
}

// typingStartEventHandler is an event handler for TypingStart events.
type typingStartEventHandler func(*Client, *EventTypingStart)

// Type returns the event type for TypingStart events.
func (eventhandler typingStartEventHandler) Type() string {
	return typingStartEventType
}

// New returns a new instance of TypingStart.
func (eventhandler typingStartEventHandler) New() interface{} {
	return &EventTypingStart{}
}

// Handle is the handler for TypingStart events.
func (eventhandler typingStartEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventTypingStart); ok {
		eventhandler(c, t)
	}
}

// userUpdateEventHandler is an event handler for UserUpdate events.
type userUpdateEventHandler func(*Client, *EventUserUpdate)

// Type returns the event type for UserUpdate events.
func (eventhandler userUpdateEventHandler) Type() string {
	return userUpdateEventType
}

// New returns a new instance of UserUpdate.
func (eventhandler userUpdateEventHandler) New() interface{} {
	return &EventUserUpdate{}
}

// Handle is the handler for UserUpdate events.
func (eventhandler userUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventUserUpdate); ok {
		eventhandler(c, t)
	}
}

// voiceServerUpdateEventHandler is an event handler for VoiceServerUpdate events.
type voiceServerUpdateEventHandler func(*Client, *EventVoiceServerUpdate)

// Type returns the event type for VoiceServerUpdate events.
func (eventhandler voiceServerUpdateEventHandler) Type() string {
	return voiceServerUpdateEventType
}

// New returns a new instance of VoiceServerUpdate.
func (eventhandler voiceServerUpdateEventHandler) New() interface{} {
	return &EventVoiceServerUpdate{}
}

// Handle is the handler for VoiceServerUpdate events.
func (eventhandler voiceServerUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventVoiceServerUpdate); ok {
		eventhandler(c, t)
	}
}

// voiceStateUpdateEventHandler is an event handler for VoiceStateUpdate events.
type voiceStateUpdateEventHandler func(*Client, *EventVoiceStateUpdate)

// Type returns the event type for VoiceStateUpdate events.
func (eventhandler voiceStateUpdateEventHandler) Type() string {
	return voiceStateUpdateEventType
}

// New returns a new instance of VoiceStateUpdate.
func (eventhandler voiceStateUpdateEventHandler) New() interface{} {
	return &EventVoiceStateUpdate{}
}

// Handle is the handler for VoiceStateUpdate events.
func (eventhandler voiceStateUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventVoiceStateUpdate); ok {
		eventhandler(c, t)
	}
}

// webhooksUpdateEventHandler is an event handler for WebhooksUpdate events.
type webhooksUpdateEventHandler func(*Client, *EventWebhooksUpdate)

// Type returns the event type for WebhooksUpdate events.
func (eventhandler webhooksUpdateEventHandler) Type() string {
	return webhooksUpdateEventType
}

// New returns a new instance of WebhooksUpdate.
func (eventhandler webhooksUpdateEventHandler) New() interface{} {
	return &EventWebhooksUpdate{}
}

// Handle is the handler for WebhooksUpdate events.
func (eventhandler webhooksUpdateEventHandler) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventWebhooksUpdate); ok {
		eventhandler(c, t)
	}
}

func handlerForInterface(handler interface{}) EventHandler {
	switch v := handler.(type) {
	case func(*Client, *EventChannelCreate):
		return channelCreateEventHandler(v)
	case func(*Client, *EventChannelDelete):
		return channelDeleteEventHandler(v)
	case func(*Client, *EventChannelPinsUpdate):
		return channelPinsUpdateEventHandler(v)
	case func(*Client, *EventChannelUpdate):
		return channelUpdateEventHandler(v)
	case func(*Client, *EventGuildBanAdd):
		return guildBanAddEventHandler(v)
	case func(*Client, *EventGuildBanRemove):
		return guildBanRemoveEventHandler(v)
	case func(*Client, *EventGuildCreate):
		return guildCreateEventHandler(v)
	case func(*Client, *EventGuildDelete):
		return guildDeleteEventHandler(v)
	case func(*Client, *EventGuildEmojisUpdate):
		return guildEmojisUpdateEventHandler(v)
	case func(*Client, *EventGuildIntegrationsUpdate):
		return guildIntegrationsUpdateEventHandler(v)
	case func(*Client, *EventGuildMemberAdd):
		return guildMemberAddEventHandler(v)
	case func(*Client, *EventGuildMemberRemove):
		return guildMemberRemoveEventHandler(v)
	case func(*Client, *EventGuildMemberUpdate):
		return guildMemberUpdateEventHandler(v)
	case func(*Client, *EventGuildRoleCreate):
		return guildRoleCreateEventHandler(v)
	case func(*Client, *EventGuildRoleDelete):
		return guildRoleDeleteEventHandler(v)
	case func(*Client, *EventGuildRoleUpdate):
		return guildRoleUpdateEventHandler(v)
	case func(*Client, *EventGuildUpdate):
		return guildUpdateEventHandler(v)
	case func(*Client, *EventHello):
		return helloEventHandler(v)
	case func(*Client, *EventInviteCreate):
		return inviteCreateEventHandler(v)
	case func(*Client, *EventInviteDelete):
		return inviteDeleteEventHandler(v)
	case func(*Client, *EventMessageCreate):
		return messageCreateEventHandler(v)
	case func(*Client, *EventMessageDelete):
		return messageDeleteEventHandler(v)
	case func(*Client, *EventMessageDeleteBulk):
		return messageDeleteBulkEventHandler(v)
	case func(*Client, *EventMessageReactionAdd):
		return messageReactionAddEventHandler(v)
	case func(*Client, *EventMessageReactionRemove):
		return messageReactionRemoveEventHandler(v)
	case func(*Client, *EventMessageReactionRemoveAll):
		return messageReactionRemoveAllEventHandler(v)
	case func(*Client, *EventMessageReactionRemoveEmoji):
		return messageReactionRemoveEmojiEventHandler(v)
	case func(*Client, *EventMessageUpdate):
		return messageUpdateEventHandler(v)
	case func(*Client, *EventPresenceUpdate):
		return presenceUpdateEventHandler(v)
	case func(*Client, *EventReady):
		return readyEventHandler(v)
	case func(*Client, *EventReconnect):
		return reconnectEventHandler(v)
	case func(*Client, *EventResumed):
		return resumedEventHandler(v)
	case func(*Client, *EventTypingStart):
		return typingStartEventHandler(v)
	case func(*Client, *EventUserUpdate):
		return userUpdateEventHandler(v)
	case func(*Client, *EventVoiceServerUpdate):
		return voiceServerUpdateEventHandler(v)
	case func(*Client, *EventVoiceStateUpdate):
		return voiceStateUpdateEventHandler(v)
	case func(*Client, *EventWebhooksUpdate):
		return webhooksUpdateEventHandler(v)
	}
	return nil
}

// EventHandler represents any EventHandler
type EventHandler interface {
	Type() string
	New() interface{}
}

var eventHandlers = map[string]EventHandler{}

func addEventHandler(eventhandler EventHandler) {
	eventHandlers[eventhandler.Type()] = eventhandler
}

// AddEventHandlers maps all event handlers
func AddEventHandlers() {
	addEventHandler(channelCreateEventHandler(nil))
	addEventHandler(channelDeleteEventHandler(nil))
	addEventHandler(channelPinsUpdateEventHandler(nil))
	addEventHandler(channelUpdateEventHandler(nil))
	addEventHandler(guildBanAddEventHandler(nil))
	addEventHandler(guildBanRemoveEventHandler(nil))
	addEventHandler(guildCreateEventHandler(nil))
	addEventHandler(guildDeleteEventHandler(nil))
	addEventHandler(guildEmojisUpdateEventHandler(nil))
	addEventHandler(guildIntegrationsUpdateEventHandler(nil))
	addEventHandler(guildMemberAddEventHandler(nil))
	addEventHandler(guildMemberRemoveEventHandler(nil))
	addEventHandler(guildMemberUpdateEventHandler(nil))
	addEventHandler(guildRoleCreateEventHandler(nil))
	addEventHandler(guildRoleDeleteEventHandler(nil))
	addEventHandler(guildRoleUpdateEventHandler(nil))
	addEventHandler(guildUpdateEventHandler(nil))
	addEventHandler(helloEventHandler(nil))
	addEventHandler(inviteCreateEventHandler(nil))
	addEventHandler(inviteDeleteEventHandler(nil))
	addEventHandler(messageCreateEventHandler(nil))
	addEventHandler(messageDeleteEventHandler(nil))
	addEventHandler(messageDeleteBulkEventHandler(nil))
	addEventHandler(messageReactionAddEventHandler(nil))
	addEventHandler(messageReactionRemoveEventHandler(nil))
	addEventHandler(messageReactionRemoveAllEventHandler(nil))
	addEventHandler(messageReactionRemoveEmojiEventHandler(nil))
	addEventHandler(messageUpdateEventHandler(nil))
	addEventHandler(presenceUpdateEventHandler(nil))
	addEventHandler(readyEventHandler(nil))
	addEventHandler(reconnectEventHandler(nil))
	addEventHandler(resumedEventHandler(nil))
	addEventHandler(typingStartEventHandler(nil))
	addEventHandler(userUpdateEventHandler(nil))
	addEventHandler(voiceServerUpdateEventHandler(nil))
	addEventHandler(voiceStateUpdateEventHandler(nil))
	addEventHandler(webhooksUpdateEventHandler(nil))
}
