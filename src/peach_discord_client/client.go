package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Client represents connection to discord.
type Client struct {

	// Logger
	Log *logrus.Logger

	// Authentification
	TOKEN string

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

	// Shard Coordinator
	ShardCoordinatorURL string

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

	// Websocket Connection
	wsConn  *websocket.Conn
	wsMutex sync.Mutex
	sync.RWMutex

	// HTTP Client
	httpClient *http.Client

	// Snowflake node to generate snowflakes
	Snowflake snowflake.Node
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

// SCGetShard sends a getshard request to the shard coordinator
func SCGetShard(c *Client) error {

	tempClient := &http.Client{}
	req, err := http.NewRequest("GET", c.ShardCoordinatorURL+"getshard", nil)
	if err != nil && err == errors.New("EOF") {
		time.Sleep(time.Second * 5)
	} else if err != nil {
		return err
	}
	resp, err := tempClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		return errors.New("Requesting ShardID failed - all shards assigned")
	}

	ShardCoordinator := ShardCoordinatorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&ShardCoordinator)
	if err != nil {
		return err
	}
	c.ShardCount = ShardCoordinator.TotalShards
	c.ShardID = ShardCoordinator.ShardID
	c.GatewayURL = ShardCoordinator.GatewayURL

	c.Log.Debugf("Websocket: Received from shard coordinator: %v", ShardCoordinator)
	return nil
}

// SCReserveShard reserves a shard
func SCReserveShard(c *Client) error {

	tempClient := &http.Client{}
	URL := c.ShardCoordinatorURL + fmt.Sprintf("reserveshard?shardid=%v", c.ShardID)
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return err
	}
	resp, err := tempClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotAcceptable {
		err := SCGetShard(c)
		if err != nil {
			return err
		}
		err = SCReserveShard(c)
		if err != nil {
			return err
		}
	} else if resp.StatusCode != http.StatusCreated {
		c.Log.Errorf("Websocket received unexpected response from shard coordinator. Expected StatusCode 200 got %v instead", resp.StatusCode)
	}
	return nil
}

// CreateClient creates a new discord client
func CreateClient(log *logrus.Logger, sharded bool) (c *Client, err error) {

	c = &Client{Sequence: new(int64), Log: log}

	// Parse shard coordinator for gateway url and shardID

	if sharded {
		// Set ShardCoordinatorURL
		c.ShardCoordinatorURL = "http://" + os.Getenv("PEACH_SHARD_COORDINATOR_SERVICE_HOST") + ":8080/api/v1/"

		err = SCGetShard(c)
		if err != nil {
			return nil, err
		}
		err = SCReserveShard(c)
		if err != nil {
			return nil, err
		}
	}

	c.httpClient = &http.Client{}

	return
}
