package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Shard shutup compiler
type Shard struct {
	ShardID          int       `json:"shard_id"`
	Reserved         bool      `json:"reserved"`
	Active           bool      `json:"active"`
	LastHeartbeat    time.Time `json:"last_heartbeat"`
	MissedHeartbeats int       `json:"missed_heartbeats"`
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

type Launcher struct {
	ID            string
	MaxClients    int
	ActiveClients int
}

type Coordinator struct {
	DB                Database
	Config            Config
	httpClient        *http.Client
	log               *logrus.Logger
	GatewayURL        string
	Bots              map[string]*Bot      `json:"bots"`
	Launchers         map[string]*Launcher `json:"launchers"`
	lock              sync.Mutex
	heartbeatInterval string
	RequiredClients   int
	ActiveClients     int
}

// Bot shutup compiler
type Bot struct {
	Username   string         `json:"username"`
	ShardCount int            `json:"shard_count"`
	Shards     map[int]*Shard `json:"shards"`
	Token      string         `json:"token"`
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
