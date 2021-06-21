package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

// Client represents connection to discord.
type Client struct {
	UserAgent string

	// Logger
	Log *logrus.Logger

	// Authentification
	TOKEN         string
	CLUSTERSECRET string

	// Settings
	Compress           bool
	LargeThreshold     int // total number of members where the gateway will stop sending offline members in the guild member list
	GuildSubscriptions bool
	Intents            int

	// Sharding
	Sharded    bool
	ShardID    int
	ShardCount int

	// Gateway URL
	GatewayURL string

	// Coordinator
	CoordinatorURL string

	// Connected represents the clients connection status
	Connected chan interface{}
	Reconnect chan interface{}
	Quit      chan interface{}

	// Session
	SessionID string
	Sequence  *int64

	// User
	User *User

	// Heartbeat
	HeartbeatInterval            time.Duration // Interval in which client should sent heartbeats
	LastHeartbeatAck             time.Time     // Last time the client received a heartbeat acknowledgement
	MissingHeartbeatAcks         time.Duration // Number of Acks that can be missed before reconnecting
	CoordinatorHeartbeatInterval string

	// Websocket Connection
	wsConn  *websocket.Conn
	wsMutex sync.Mutex
	sync.RWMutex

	// HTTP
	Ratelimiter *Ratelimiter
	httpClient  *http.Client
	httpRetries int

	// Cache
	Guilds       *[]Guild
	GuildCache   *cache.Cache
	ChannelCache *cache.Cache
	Settings     map[string]cfgSettings // Map guildIDs to settings and cache that shit

	// Extensions
	Extensions Extensions

	// Starttime
	Starttime time.Time
}

// Run starts various background routines and starts listeners
func (c *Client) Run() error {
	c.Log.Info("Starting Websocket")

	err := c.CreateWebsocket()
	if err != nil {
		return err
	}

	return nil
}

// CreateClient creates a new discord client
func CreateClient(log *logrus.Logger, sharded bool, coordiantorURL string, secret string) (c *Client, spotifyID *string, spotifySecret *string, err error) {

	c = &Client{Sequence: new(int64), Log: log}
	c.Starttime = time.Now()
	c.Settings = make(map[string]cfgSettings)
	c.Ratelimiter = CreateRatelimiter()

	// login to coordinator

	if sharded {
		c.CoordinatorURL = coordiantorURL
		c.CLUSTERSECRET = secret

		spotifyID, spotifySecret, err = c.CoordinatorLogin()
		if err != nil {
			return nil, nil, nil, err
		}
	}

	c.Reconnect = make(chan interface{})
	c.Quit = make(chan interface{})

	return
}

// FetchAll retrieves all Guild settings from the database
func (c *Client) FetchAll() (err error) {
	if c.Guilds == nil {
		c.Guilds, err = c.GetUserGuilds()
		if err != nil {
			return err
		}

	}

	for _, guild := range *c.Guilds {
		err = c.getGuildSettings(guild.ID)

		if err != nil {
			c.Log.Error(err)
		}
	}

	return nil

}
