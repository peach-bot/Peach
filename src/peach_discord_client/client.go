package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

	// Client Coordinator
	ClientCoordinatorURL string

	// Connected represents the clients connection status
	Connected chan interface{}
	Reconnect chan interface{}
	Quit      chan interface{}

	// Session
	SessionID string
	Sequence  *int64

	// User
	User User

	// Heartbeat
	HeartbeatInterval    time.Duration // Interval in which client should sent heartbeats
	LastHeartbeatAck     time.Time     // Last time the client received a heartbeat acknowledgement
	MissingHeartbeatAcks time.Duration // Number of Acks that can be missed before reconnecting
	CCHeartbeatInterval  string

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

// CCLogin registeres the client in the client coordinator and reserves a shard
func CCLogin(c *Client) error {

	tempClient := &http.Client{}
	req, err := http.NewRequest("GET", c.ClientCoordinatorURL+"login", nil)
	if err != nil && err == errors.New("EOF") {
		time.Sleep(time.Second * 5)
		return CCLogin(c)
	} else if err != nil {
		return err
	}
	req.Header.Add("authorization", c.CLUSTERSECRET)
	resp, err := tempClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return errors.New("Requesting ShardID failed - all shards assigned")
	} else if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("ClientCoordinator sent unexpected response: Want 200Ok Got %s", resp.Status))
	}

	ClientCoordinator := ClientCoordinatorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&ClientCoordinator)
	if err != nil {
		return err
	}
	c.TOKEN = ClientCoordinator.Token
	c.ShardCount = ClientCoordinator.TotalShards
	c.ShardID = ClientCoordinator.ShardID
	c.GatewayURL = ClientCoordinator.GatewayURL
	c.CCHeartbeatInterval = ClientCoordinator.HeartbeatInterval
	c.Log.Debugf("Websocket: Received from client coordinator: %v", ClientCoordinator)
	return nil
}

func setCCRequestHeaders(c *Client, req *http.Request) *http.Request {
	req.Header.Add("authorization", c.CLUSTERSECRET)
	req.Header.Add("bot_id", c.User.ID)
	req.Header.Add("shard_id", strconv.Itoa(c.ShardID))
	req.Close = true
	return req
}

// CCReady asdasd
func CCReady(c *Client) error {
	tempClient := &http.Client{}
	req, err := http.NewRequest("GET", c.ClientCoordinatorURL+"ready", nil)
	if err != nil {
		return err
	}
	req = setCCRequestHeaders(c, req)
	resp, err := tempClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.Log.Fatalf("Websocket received unexpected response from client coordinator. Expected Status 200 OK got %s instead", resp.Status)
	}
	return nil
}

// CCHeartbeat stfu
func (c *Client) CCHeartbeat() {
	c.Log.Info("Started sending heartbeat to client coordinator.")
	interval, err := time.ParseDuration(c.CCHeartbeatInterval)
	if err != nil {
		c.Log.Fatal(err)
	}
	ticker := time.NewTicker(interval)
	for {
		tempClient := &http.Client{}
		req, err := http.NewRequest("GET", c.ClientCoordinatorURL+"heartbeat", nil)
		if err != nil {
			c.Log.Error(err)
		}
		req = setCCRequestHeaders(c, req)
		resp, err := tempClient.Do(req)
		if err != nil {
			c.Log.Error(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			c.Log.Errorf("Websocket received unexpected response from client coordinator. Expected Status 200 OK got %s instead", resp.Status)
		}
		select {
		case <-ticker.C:
		}
	}
}

// CreateClient creates a new discord client
func CreateClient(log *logrus.Logger, sharded bool, ccURL string, secret string) (c *Client, err error) {

	c = &Client{Sequence: new(int64), Log: log}
	c.Starttime = time.Now()
	c.Settings = make(map[string]cfgSettings)
	c.Ratelimiter = CreateRatelimiter()

	// Parse client coordinator for gateway url and shardID

	if sharded {
		c.ClientCoordinatorURL = ccURL
		c.CLUSTERSECRET = secret

		err = CCLogin(c)
		if err != nil {
			return nil, err
		}
	}

	c.Reconnect = make(chan interface{})
	c.Quit = make(chan interface{})

	return
}

// GetChannel retrieves the Channel object for a given ID
func (c *Client) GetChannel(ID string) (ch *Channel, err error) {

	cachedChannel, cached := c.ChannelCache.Get(ID)

	if cached {
		channel := cachedChannel.(Channel)
		ch = &channel
	} else {
		ch, err = c.getChannel(ID)
		return
	}

	return
}

// GetGuild retrieves the Guild object for a given ID
func (c *Client) GetGuild(ID string) (g *Guild, err error) {

	cachedGuild, cached := c.GuildCache.Get(ID)

	if cached {
		guild := cachedGuild.(Guild)
		g = &guild
	} else {
		g, err = c.getGuild(ID)
		return
	}

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
