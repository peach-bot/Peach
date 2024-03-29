package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errInvalidToken = errors.New("passed invalid token")

func (c *Coordinator) getGatewayBot(token string) (*getgatewayresponse, error) {
	var gwr getgatewayresponse
	req, err := http.NewRequest("GET", "https://discord.com/api/v8/gateway/bot", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bot %v", token))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&gwr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errInvalidToken
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received unexpected status from discord: expected 200 OK got %s instead", resp.Status)
	}
	return &gwr, nil
}

func (c *Coordinator) create() error {
	c.httpClient = new(http.Client)
	c.gettokens()
	c.Bots = make(map[string]*Bot)

	for _, token := range tokens {
		gwr, err := c.getGatewayBot(token)
		if err != nil {
			return err
		}
		user, err := c.getUser(token)
		if err != nil {
			return err
		}
		c.GatewayURL = gwr.URL
		c.Bots[user.ID] = &Bot{}
		bot := c.Bots[user.ID]
		bot.Username = user.Username + "#" + user.Discriminator
		bot.ShardCount = gwr.Shards
		bot.Token = token
		bot.Shards = make(map[int]*Shard)
		for i := 0; i < gwr.Shards; i++ {
			bot.Shards[i] = &Shard{i, false, false, time.Now(), 0}
			c.RequiredClients++
		}
	}
	return nil
}

func (c *Coordinator) getUser(token string) (*User, error) {
	var user User
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bot %v", token))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received unexpected status from discord: expected 200 OK got %s instead", resp.Status)
	}
	return &user, nil
}

func (c *Coordinator) nextShard() (*Bot, *Shard) {
	for _, bot := range c.Bots {
		for _, shard := range bot.Shards {
			if !shard.Reserved {
				return bot, shard
			}
		}
	}
	return nil, nil
}

func (c *Coordinator) shardManager(bot *Bot, shard *Shard) {
	interval, err := time.ParseDuration(c.heartbeatInterval)
	if err != nil {
		c.log.Fatal(err)
	}
	time.Sleep((interval) - time.Since(shard.LastHeartbeat))
	ticker := time.NewTicker(interval)
	shard.MissedHeartbeats = 0
	defer ticker.Stop()
	for {
		if shard.Reserved && !shard.Active {
			shard.Reserved = false
			return
		}
		if shard.Reserved && shard.Active {
			shard.MissedHeartbeats++
		}
		if shard.MissedHeartbeats > 3 {
			shard.Reserved, shard.Active, shard.MissedHeartbeats = false, false, 0
			return
		}

		select {
		case <-ticker.C:
		}
	}
}
