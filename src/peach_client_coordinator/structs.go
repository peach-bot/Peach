package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Shard shutup compiler
type Shard struct {
	ShardID          int
	Reserved         bool
	Active           bool
	LastHeartbeat    time.Time
	MissedHeartbeats int
}

type getgatewayresponse struct {
	URL               string
	Shards            int
	SessionStartLimit struct {
		Total      int
		Remaining  int
		ResetAfter int
	}
}

type clientCoordinator struct {
	httpClient *http.Client
	log        *logrus.Logger
	GatewayURL string
	Bots       map[string]*Bot
}

// Bot shutup compiler
type Bot struct {
	Shards     map[int]*Shard
	ShardCount int
	Token      string
}

// User represents a discord user
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot,omitempty"`
	System        bool   `json:"system,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Language      string `json:"locale,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Email         string `json:"email,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	NitroType     int    `json:"premium_type,omitempty"`
}
