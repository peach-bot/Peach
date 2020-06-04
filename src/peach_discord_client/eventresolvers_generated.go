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
	resumeEventType                  = "RESUME"
	webhooksUpdateEventType          = "WEBHOOKS_UPDATE"
)

// channelCreateEventResolver is an event resolver for ChannelCreate events.
type channelCreateEventResolver func(*Client, *EventChannelCreate)

// Type returns the event type for ChannelCreate events.
func (eventresolver channelCreateEventResolver) Type() string {
	return channelCreateEventType
}

// New returns a new instance of ChannelCreate.
func (eventresolver channelCreateEventResolver) New() interface{} {
	return &EventChannelCreate{}
}

// Handle is the handler for ChannelCreate events.
func (eventresolver channelCreateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelCreate); ok {
		eventresolver(c, t)
	}
}

// channelDeleteEventResolver is an event resolver for ChannelDelete events.
type channelDeleteEventResolver func(*Client, *EventChannelDelete)

// Type returns the event type for ChannelDelete events.
func (eventresolver channelDeleteEventResolver) Type() string {
	return channelDeleteEventType
}

// New returns a new instance of ChannelDelete.
func (eventresolver channelDeleteEventResolver) New() interface{} {
	return &EventChannelDelete{}
}

// Handle is the handler for ChannelDelete events.
func (eventresolver channelDeleteEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelDelete); ok {
		eventresolver(c, t)
	}
}

// channelPinsUpdateEventResolver is an event resolver for ChannelPinsUpdate events.
type channelPinsUpdateEventResolver func(*Client, *EventChannelPinsUpdate)

// Type returns the event type for ChannelPinsUpdate events.
func (eventresolver channelPinsUpdateEventResolver) Type() string {
	return channelPinsUpdateEventType
}

// New returns a new instance of ChannelPinsUpdate.
func (eventresolver channelPinsUpdateEventResolver) New() interface{} {
	return &EventChannelPinsUpdate{}
}

// Handle is the handler for ChannelPinsUpdate events.
func (eventresolver channelPinsUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelPinsUpdate); ok {
		eventresolver(c, t)
	}
}

// channelUpdateEventResolver is an event resolver for ChannelUpdate events.
type channelUpdateEventResolver func(*Client, *EventChannelUpdate)

// Type returns the event type for ChannelUpdate events.
func (eventresolver channelUpdateEventResolver) Type() string {
	return channelUpdateEventType
}

// New returns a new instance of ChannelUpdate.
func (eventresolver channelUpdateEventResolver) New() interface{} {
	return &EventChannelUpdate{}
}

// Handle is the handler for ChannelUpdate events.
func (eventresolver channelUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventChannelUpdate); ok {
		eventresolver(c, t)
	}
}

// guildBanAddEventResolver is an event resolver for GuildBanAdd events.
type guildBanAddEventResolver func(*Client, *EventGuildBanAdd)

// Type returns the event type for GuildBanAdd events.
func (eventresolver guildBanAddEventResolver) Type() string {
	return guildBanAddEventType
}

// New returns a new instance of GuildBanAdd.
func (eventresolver guildBanAddEventResolver) New() interface{} {
	return &EventGuildBanAdd{}
}

// Handle is the handler for GuildBanAdd events.
func (eventresolver guildBanAddEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildBanAdd); ok {
		eventresolver(c, t)
	}
}

// guildBanRemoveEventResolver is an event resolver for GuildBanRemove events.
type guildBanRemoveEventResolver func(*Client, *EventGuildBanRemove)

// Type returns the event type for GuildBanRemove events.
func (eventresolver guildBanRemoveEventResolver) Type() string {
	return guildBanRemoveEventType
}

// New returns a new instance of GuildBanRemove.
func (eventresolver guildBanRemoveEventResolver) New() interface{} {
	return &EventGuildBanRemove{}
}

// Handle is the handler for GuildBanRemove events.
func (eventresolver guildBanRemoveEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildBanRemove); ok {
		eventresolver(c, t)
	}
}

// guildCreateEventResolver is an event resolver for GuildCreate events.
type guildCreateEventResolver func(*Client, *EventGuildCreate)

// Type returns the event type for GuildCreate events.
func (eventresolver guildCreateEventResolver) Type() string {
	return guildCreateEventType
}

// New returns a new instance of GuildCreate.
func (eventresolver guildCreateEventResolver) New() interface{} {
	return &EventGuildCreate{}
}

// Handle is the handler for GuildCreate events.
func (eventresolver guildCreateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildCreate); ok {
		eventresolver(c, t)
	}
}

// guildDeleteEventResolver is an event resolver for GuildDelete events.
type guildDeleteEventResolver func(*Client, *EventGuildDelete)

// Type returns the event type for GuildDelete events.
func (eventresolver guildDeleteEventResolver) Type() string {
	return guildDeleteEventType
}

// New returns a new instance of GuildDelete.
func (eventresolver guildDeleteEventResolver) New() interface{} {
	return &EventGuildDelete{}
}

// Handle is the handler for GuildDelete events.
func (eventresolver guildDeleteEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildDelete); ok {
		eventresolver(c, t)
	}
}

// guildEmojisUpdateEventResolver is an event resolver for GuildEmojisUpdate events.
type guildEmojisUpdateEventResolver func(*Client, *EventGuildEmojisUpdate)

// Type returns the event type for GuildEmojisUpdate events.
func (eventresolver guildEmojisUpdateEventResolver) Type() string {
	return guildEmojisUpdateEventType
}

// New returns a new instance of GuildEmojisUpdate.
func (eventresolver guildEmojisUpdateEventResolver) New() interface{} {
	return &EventGuildEmojisUpdate{}
}

// Handle is the handler for GuildEmojisUpdate events.
func (eventresolver guildEmojisUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildEmojisUpdate); ok {
		eventresolver(c, t)
	}
}

// guildIntegrationsUpdateEventResolver is an event resolver for GuildIntegrationsUpdate events.
type guildIntegrationsUpdateEventResolver func(*Client, *EventGuildIntegrationsUpdate)

// Type returns the event type for GuildIntegrationsUpdate events.
func (eventresolver guildIntegrationsUpdateEventResolver) Type() string {
	return guildIntegrationsUpdateEventType
}

// New returns a new instance of GuildIntegrationsUpdate.
func (eventresolver guildIntegrationsUpdateEventResolver) New() interface{} {
	return &EventGuildIntegrationsUpdate{}
}

// Handle is the handler for GuildIntegrationsUpdate events.
func (eventresolver guildIntegrationsUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildIntegrationsUpdate); ok {
		eventresolver(c, t)
	}
}

// guildMemberAddEventResolver is an event resolver for GuildMemberAdd events.
type guildMemberAddEventResolver func(*Client, *EventGuildMemberAdd)

// Type returns the event type for GuildMemberAdd events.
func (eventresolver guildMemberAddEventResolver) Type() string {
	return guildMemberAddEventType
}

// New returns a new instance of GuildMemberAdd.
func (eventresolver guildMemberAddEventResolver) New() interface{} {
	return &EventGuildMemberAdd{}
}

// Handle is the handler for GuildMemberAdd events.
func (eventresolver guildMemberAddEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildMemberAdd); ok {
		eventresolver(c, t)
	}
}

// guildMemberRemoveEventResolver is an event resolver for GuildMemberRemove events.
type guildMemberRemoveEventResolver func(*Client, *EventGuildMemberRemove)

// Type returns the event type for GuildMemberRemove events.
func (eventresolver guildMemberRemoveEventResolver) Type() string {
	return guildMemberRemoveEventType
}

// New returns a new instance of GuildMemberRemove.
func (eventresolver guildMemberRemoveEventResolver) New() interface{} {
	return &EventGuildMemberRemove{}
}

// Handle is the handler for GuildMemberRemove events.
func (eventresolver guildMemberRemoveEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildMemberRemove); ok {
		eventresolver(c, t)
	}
}

// guildMemberUpdateEventResolver is an event resolver for GuildMemberUpdate events.
type guildMemberUpdateEventResolver func(*Client, *EventGuildMemberUpdate)

// Type returns the event type for GuildMemberUpdate events.
func (eventresolver guildMemberUpdateEventResolver) Type() string {
	return guildMemberUpdateEventType
}

// New returns a new instance of GuildMemberUpdate.
func (eventresolver guildMemberUpdateEventResolver) New() interface{} {
	return &EventGuildMemberUpdate{}
}

// Handle is the handler for GuildMemberUpdate events.
func (eventresolver guildMemberUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildMemberUpdate); ok {
		eventresolver(c, t)
	}
}

// guildRoleCreateEventResolver is an event resolver for GuildRoleCreate events.
type guildRoleCreateEventResolver func(*Client, *EventGuildRoleCreate)

// Type returns the event type for GuildRoleCreate events.
func (eventresolver guildRoleCreateEventResolver) Type() string {
	return guildRoleCreateEventType
}

// New returns a new instance of GuildRoleCreate.
func (eventresolver guildRoleCreateEventResolver) New() interface{} {
	return &EventGuildRoleCreate{}
}

// Handle is the handler for GuildRoleCreate events.
func (eventresolver guildRoleCreateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildRoleCreate); ok {
		eventresolver(c, t)
	}
}

// guildRoleDeleteEventResolver is an event resolver for GuildRoleDelete events.
type guildRoleDeleteEventResolver func(*Client, *EventGuildRoleDelete)

// Type returns the event type for GuildRoleDelete events.
func (eventresolver guildRoleDeleteEventResolver) Type() string {
	return guildRoleDeleteEventType
}

// New returns a new instance of GuildRoleDelete.
func (eventresolver guildRoleDeleteEventResolver) New() interface{} {
	return &EventGuildRoleDelete{}
}

// Handle is the handler for GuildRoleDelete events.
func (eventresolver guildRoleDeleteEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildRoleDelete); ok {
		eventresolver(c, t)
	}
}

// guildRoleUpdateEventResolver is an event resolver for GuildRoleUpdate events.
type guildRoleUpdateEventResolver func(*Client, *EventGuildRoleUpdate)

// Type returns the event type for GuildRoleUpdate events.
func (eventresolver guildRoleUpdateEventResolver) Type() string {
	return guildRoleUpdateEventType
}

// New returns a new instance of GuildRoleUpdate.
func (eventresolver guildRoleUpdateEventResolver) New() interface{} {
	return &EventGuildRoleUpdate{}
}

// Handle is the handler for GuildRoleUpdate events.
func (eventresolver guildRoleUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildRoleUpdate); ok {
		eventresolver(c, t)
	}
}

// guildUpdateEventResolver is an event resolver for GuildUpdate events.
type guildUpdateEventResolver func(*Client, *EventGuildUpdate)

// Type returns the event type for GuildUpdate events.
func (eventresolver guildUpdateEventResolver) Type() string {
	return guildUpdateEventType
}

// New returns a new instance of GuildUpdate.
func (eventresolver guildUpdateEventResolver) New() interface{} {
	return &EventGuildUpdate{}
}

// Handle is the handler for GuildUpdate events.
func (eventresolver guildUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventGuildUpdate); ok {
		eventresolver(c, t)
	}
}

// helloEventResolver is an event resolver for Hello events.
type helloEventResolver func(*Client, *EventHello)

// Type returns the event type for Hello events.
func (eventresolver helloEventResolver) Type() string {
	return helloEventType
}

// New returns a new instance of Hello.
func (eventresolver helloEventResolver) New() interface{} {
	return &EventHello{}
}

// Handle is the handler for Hello events.
func (eventresolver helloEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventHello); ok {
		eventresolver(c, t)
	}
}

// messageCreateEventResolver is an event resolver for MessageCreate events.
type messageCreateEventResolver func(*Client, *EventMessageCreate)

// Type returns the event type for MessageCreate events.
func (eventresolver messageCreateEventResolver) Type() string {
	return messageCreateEventType
}

// New returns a new instance of MessageCreate.
func (eventresolver messageCreateEventResolver) New() interface{} {
	return &EventMessageCreate{}
}

// Handle is the handler for MessageCreate events.
func (eventresolver messageCreateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventMessageCreate); ok {
		eventresolver(c, t)
	}
}

// presenceUpdateEventResolver is an event resolver for PresenceUpdate events.
type presenceUpdateEventResolver func(*Client, *EventPresenceUpdate)

// Type returns the event type for PresenceUpdate events.
func (eventresolver presenceUpdateEventResolver) Type() string {
	return presenceUpdateEventType
}

// New returns a new instance of PresenceUpdate.
func (eventresolver presenceUpdateEventResolver) New() interface{} {
	return &EventPresenceUpdate{}
}

// Handle is the handler for PresenceUpdate events.
func (eventresolver presenceUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventPresenceUpdate); ok {
		eventresolver(c, t)
	}
}

// readyEventResolver is an event resolver for Ready events.
type readyEventResolver func(*Client, *EventReady)

// Type returns the event type for Ready events.
func (eventresolver readyEventResolver) Type() string {
	return readyEventType
}

// New returns a new instance of Ready.
func (eventresolver readyEventResolver) New() interface{} {
	return &EventReady{}
}

// Handle is the handler for Ready events.
func (eventresolver readyEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventReady); ok {
		eventresolver(c, t)
	}
}

// resumeEventResolver is an event resolver for Resume events.
type resumeEventResolver func(*Client, *EventResume)

// Type returns the event type for Resume events.
func (eventresolver resumeEventResolver) Type() string {
	return resumeEventType
}

// New returns a new instance of Resume.
func (eventresolver resumeEventResolver) New() interface{} {
	return &EventResume{}
}

// Handle is the handler for Resume events.
func (eventresolver resumeEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventResume); ok {
		eventresolver(c, t)
	}
}

// webhooksUpdateEventResolver is an event resolver for WebhooksUpdate events.
type webhooksUpdateEventResolver func(*Client, *EventWebhooksUpdate)

// Type returns the event type for WebhooksUpdate events.
func (eventresolver webhooksUpdateEventResolver) Type() string {
	return webhooksUpdateEventType
}

// New returns a new instance of WebhooksUpdate.
func (eventresolver webhooksUpdateEventResolver) New() interface{} {
	return &EventWebhooksUpdate{}
}

// Handle is the handler for WebhooksUpdate events.
func (eventresolver webhooksUpdateEventResolver) Handle(c *Client, i interface{}) {
	if t, ok := i.(*EventWebhooksUpdate); ok {
		eventresolver(c, t)
	}
}

func handlerForInterface(resolver interface{}) EventResolver {
	switch v := resolver.(type) {
	case func(*Client, *EventChannelCreate):
		return channelCreateEventResolver(v)
	case func(*Client, *EventChannelDelete):
		return channelDeleteEventResolver(v)
	case func(*Client, *EventChannelPinsUpdate):
		return channelPinsUpdateEventResolver(v)
	case func(*Client, *EventChannelUpdate):
		return channelUpdateEventResolver(v)
	case func(*Client, *EventGuildBanAdd):
		return guildBanAddEventResolver(v)
	case func(*Client, *EventGuildBanRemove):
		return guildBanRemoveEventResolver(v)
	case func(*Client, *EventGuildCreate):
		return guildCreateEventResolver(v)
	case func(*Client, *EventGuildDelete):
		return guildDeleteEventResolver(v)
	case func(*Client, *EventGuildEmojisUpdate):
		return guildEmojisUpdateEventResolver(v)
	case func(*Client, *EventGuildIntegrationsUpdate):
		return guildIntegrationsUpdateEventResolver(v)
	case func(*Client, *EventGuildMemberAdd):
		return guildMemberAddEventResolver(v)
	case func(*Client, *EventGuildMemberRemove):
		return guildMemberRemoveEventResolver(v)
	case func(*Client, *EventGuildMemberUpdate):
		return guildMemberUpdateEventResolver(v)
	case func(*Client, *EventGuildRoleCreate):
		return guildRoleCreateEventResolver(v)
	case func(*Client, *EventGuildRoleDelete):
		return guildRoleDeleteEventResolver(v)
	case func(*Client, *EventGuildRoleUpdate):
		return guildRoleUpdateEventResolver(v)
	case func(*Client, *EventGuildUpdate):
		return guildUpdateEventResolver(v)
	case func(*Client, *EventHello):
		return helloEventResolver(v)
	case func(*Client, *EventMessageCreate):
		return messageCreateEventResolver(v)
	case func(*Client, *EventPresenceUpdate):
		return presenceUpdateEventResolver(v)
	case func(*Client, *EventReady):
		return readyEventResolver(v)
	case func(*Client, *EventResume):
		return resumeEventResolver(v)
	case func(*Client, *EventWebhooksUpdate):
		return webhooksUpdateEventResolver(v)
	}
	return nil
}

// EventResolver represents any EventResolver
type EventResolver interface {
	Type() string
	New() interface{}
}

var eventResolvers = map[string]EventResolver{}

func addEventResolver(eventresolver EventResolver) {
	eventResolvers[eventresolver.Type()] = eventresolver
}

// AddEventResolvers maps all event resolvers
func AddEventResolvers() {
	addEventResolver(channelCreateEventResolver(nil))
	addEventResolver(channelDeleteEventResolver(nil))
	addEventResolver(channelPinsUpdateEventResolver(nil))
	addEventResolver(channelUpdateEventResolver(nil))
	addEventResolver(guildBanAddEventResolver(nil))
	addEventResolver(guildBanRemoveEventResolver(nil))
	addEventResolver(guildCreateEventResolver(nil))
	addEventResolver(guildDeleteEventResolver(nil))
	addEventResolver(guildEmojisUpdateEventResolver(nil))
	addEventResolver(guildIntegrationsUpdateEventResolver(nil))
	addEventResolver(guildMemberAddEventResolver(nil))
	addEventResolver(guildMemberRemoveEventResolver(nil))
	addEventResolver(guildMemberUpdateEventResolver(nil))
	addEventResolver(guildRoleCreateEventResolver(nil))
	addEventResolver(guildRoleDeleteEventResolver(nil))
	addEventResolver(guildRoleUpdateEventResolver(nil))
	addEventResolver(guildUpdateEventResolver(nil))
	addEventResolver(helloEventResolver(nil))
	addEventResolver(messageCreateEventResolver(nil))
	addEventResolver(presenceUpdateEventResolver(nil))
	addEventResolver(readyEventResolver(nil))
	addEventResolver(resumeEventResolver(nil))
	addEventResolver(webhooksUpdateEventResolver(nil))
}
