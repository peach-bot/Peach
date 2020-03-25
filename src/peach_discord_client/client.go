package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

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
func CreateClient(log *logrus.Logger) (c *Client, err error) {

	c = &Client{Sequence: new(int64), Log: log}

	// Parse shard coordinator for gateway url and shardID

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

	return
}
