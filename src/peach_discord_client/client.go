package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// Run starts various background routines and starts listeners
func (c *Client) Run() error {
	c.Log.Info("Starting Websocket")

	err := c.CreateWebsocket()
	if err != nil {
		return err
	}

	return nil
}

// SCGetShards sends a getshard request to the shard coordinator
func SCGetShards(c *Client) {

	temp := &http.Client{}
	req, err := http.NewRequest("GET", c.ShardCoordinatorURL+"/api/v1/getshard", nil)
	if err != nil {
		c.Log.Error(err)
	}
	resp, err := temp.Do(req)
	if err != nil {
		c.Log.Error(err)
	}

	c.ShardCoordinator = ShardCoordinatorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&c.ShardCoordinator)
	if err != nil {
		c.Log.Error(err)
	}

	c.Log.Debugf("Websocket: Received from shard coordinator: %v", c.ShardCoordinator)
}

// CreateClient creates a new discord client
func CreateClient(log *logrus.Logger) (c *Client, err error) {

	c = &Client{Sequence: new(int64), Log: log}

	// Parse shard coordinator for gateway url and shardID
	c.GatewayURL = "wss://gateway.discord.gg/"
	c.GatewayURL = c.GatewayURL + "?v=" + APIVersion + "&encoding=json"

	// Set ShardCoordinatorURL
	c.ShardCoordinatorURL = "http://" + os.Getenv("PEACH_SHARD_COORDINATOR_SERVICE_HOST") + ":8080"

	SCGetShards(c)

	return
}
