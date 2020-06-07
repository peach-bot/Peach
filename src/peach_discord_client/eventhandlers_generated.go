package main

// Event types used to match values sent by Discord
const (
	channelCreateEventType           = "CHANNEL_CREATE"
	channelDeleteEventType           = "CHANNEL_DELETE"
	channelPinsUpdateEventType       = "CHANNEL_PINS_UPDATE"
	channelUpdateEventType           = "CHANNEL_UPDATE"
	guildBanAddEventType             = "GUILD_BAN_ADD"
	guildBanRemoveEventType          = "GUILD_BAN_REMOVE"
	guildCreateEventType             = "GUILD_CREATE"
	guildDeleteEventType             = "GUILD_DELETE"
	guildEmojisUpdateEventType       = "GUILD_EMOJIS_UPDATE"
	guildIntegrationsUpdateEventType = "GUILD_INTEGRATIONS_UPDATE"
	guildMemberAddEventType          = "GUILD_MEMBER_ADD"
	guildMemberRemoveEventType       = "GUILD_MEMBER_REMOVE"
	guildMemberUpdateEventType       = "GUILD_MEMBER_UPDATE"
	guildRoleCreateEventType         = "GUILD_ROLE_CREATE"
	guildRoleDeleteEventType         = "GUILD_ROLE_DELETE"
	guildRoleUpdateEventType         = "GUILD_ROLE_UPDATE"
	guildUpdateEventType             = "GUILD_UPDATE"
	helloEventType                   = "HELLO"
	messageCreateEventType           = "MESSAGE_CREATE"
	presenceUpdateEventType          = "PRESENCE_UPDATE"
	readyEventType                   = "READY"
	resumedEventType                 = "RESUMED"
	webhooksUpdateEventType          = "WEBHOOKS_UPDATE"
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
	case func(*Client, *EventMessageCreate):
		return messageCreateEventHandler(v)
	case func(*Client, *EventPresenceUpdate):
		return presenceUpdateEventHandler(v)
	case func(*Client, *EventReady):
		return readyEventHandler(v)
	case func(*Client, *EventResumed):
		return resumedEventHandler(v)
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
	addEventHandler(messageCreateEventHandler(nil))
	addEventHandler(presenceUpdateEventHandler(nil))
	addEventHandler(readyEventHandler(nil))
	addEventHandler(resumedEventHandler(nil))
	addEventHandler(webhooksUpdateEventHandler(nil))
}
