package main

// Event types used to match values sent by Discord
const (
	helloEventType          = "HELLO"
	messageCreateEventType  = "MESSAGE_CREATE"
	presenceUpdateEventType = "PRESENCE_UPDATE"
	readyEventType          = "READY"
)

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

func handlerForInterface(resolver interface{}) EventResolver {
	switch v := resolver.(type) {
	case func(*Client, *EventHello):
		return helloEventResolver(v)
	case func(*Client, *EventMessageCreate):
		return messageCreateEventResolver(v)
	case func(*Client, *EventPresenceUpdate):
		return presenceUpdateEventResolver(v)
	case func(*Client, *EventReady):
		return readyEventResolver(v)
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
	addEventResolver(helloEventResolver(nil))
	addEventResolver(messageCreateEventResolver(nil))
	addEventResolver(presenceUpdateEventResolver(nil))
	addEventResolver(readyEventResolver(nil))
}
