package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// CoordinatorLogin registeres the client in the coordinator and reserves a shard
func (c *Client) CoordinatorLogin() error {

	tempClient := &http.Client{}
	req, err := http.NewRequest("GET", c.CoordinatorURL+"login", nil)
	if err != nil && err == errors.New("EOF") {
		time.Sleep(time.Second * 5)
		return c.CoordinatorLogin()
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

	coordresp := CoordinatorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&coordresp)
	if err != nil {
		return err
	}
	c.ShardCount = coordresp.TotalShards
	c.ShardID = coordresp.ShardID
	c.GatewayURL = coordresp.GatewayURL
	c.CoordinatorHeartbeatInterval = coordresp.HeartbeatInterval
	c.TOKEN = coordresp.Token

	if redactSensitive {
		coordresp.Token = "[REDACTED]"
	}

	c.Log.Debugf("Websocket: Received from coordinator: %v", coordresp)
	return nil
}

func (c *Client) setCoordinatorRequestHeaders(req *http.Request) *http.Request {
	req.Header.Add("authorization", c.CLUSTERSECRET)
	req.Header.Add("bot_id", c.User.ID)
	req.Header.Add("shard_id", strconv.Itoa(c.ShardID))
	req.Close = true
	return req
}

// CCReady asdasd
func (c *Client) CoordiantorReady() error {
	tempClient := &http.Client{}
	req, err := http.NewRequest("GET", c.CoordinatorURL+"ready", nil)
	if err != nil {
		return err
	}
	req = c.setCoordinatorRequestHeaders(req)
	resp, err := tempClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.Log.Fatalf("Websocket received unexpected response from coordinator. Expected Status 200 OK got %s instead", resp.Status)
	}
	return nil
}

// CCHeartbeat stfu
func (c *Client) CoordinatorHeartbeat() {
	c.Log.Info("Started sending heartbeat to coordinator.")
	interval, err := time.ParseDuration(c.CoordinatorHeartbeatInterval)
	if err != nil {
		c.Log.Fatal(err)
	}
	ticker := time.NewTicker(interval)
	for {
		tempClient := &http.Client{}
		req, err := http.NewRequest("GET", c.CoordinatorURL+"heartbeat", nil)
		if err != nil {
			c.Log.Error(err)
		}
		req = c.setCoordinatorRequestHeaders(req)
		resp, err := tempClient.Do(req)
		if err != nil {
			c.Log.Error(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			c.Log.Errorf("Websocket received unexpected response from coordinator. Expected Status 200 OK got %s instead", resp.Status)
		}
		select {
		case <-ticker.C:
		}
	}
}
