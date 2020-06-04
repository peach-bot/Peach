package main

// Emoji represents a discord emoji
type Emoji struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Roles         []*Role `json:"roles,omitempty"`
	User          User    `json:"user,omitempty"`
	RequireColons bool    `json:"require_colons,omitempty"`
	Managed       bool    `json:"managed,omitempty"`
	Animated      bool    `json:"animated,omitempty"`
	Available     bool    `json:"available,omitempty"`
}
