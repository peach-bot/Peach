package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (c *clientCoordinator) verifyAuth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	authsecret := r.Header.Get("authorization")
	if authsecret != clustersecret {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New(("Auth code didn't match cluster secret"))
	}
	return nil
}

func (c *clientCoordinator) verifyBotShard(w http.ResponseWriter, r *http.Request) (*Bot, *Shard, error) {

	botID := r.Header.Get("bot_id")
	if botID == "" {
		return nil, nil, errors.New("Header missing bot_id")
	}

	sid := r.Header.Get("shard_id")
	if sid == "" {
		return nil, nil, errors.New("Header missing shard_id")
	}

	shardID, err := strconv.Atoi(sid)
	if err != nil {
		return nil, nil, errors.New("invalid shard_id")
	}

	bot := c.Bots[botID]
	shard := bot.Shards[shardID]
	if bot == nil {
		return nil, nil, errors.New("invalid bot_id")
	}
	if shard == nil {
		return nil, nil, errors.New("invalid shard_id")
	}

	return bot, shard, nil
}

func (c *clientCoordinator) pathLogin(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		return
	}

	c.lock.Lock()
	bot, shard := c.nextShard()
	if bot == nil || shard == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := fmt.Sprintf(`{"token": "%s", "total_shards": %d, "assigned_shard": %d, "gateway_url": "%s"}`, bot.Token, bot.ShardCount, shard.ShardID, c.GatewayURL)
	shard.Reserved = true
	c.lock.Unlock()
	shard.LastHeartbeat = time.Now()
	go c.shardManager(bot, shard)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (c *clientCoordinator) pathReady(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		return
	}
	_, shard, err := c.verifyBotShard(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	if shard.Active {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Shard already active"))
		return
	}
	shard.Active = true
	w.WriteHeader(http.StatusOK)
}

func (c *clientCoordinator) pathHeartbeat(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		return
	}
	_, shard, err := c.verifyBotShard(w, r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	shard.LastHeartbeat = time.Now()
	shard.MissedHeartbeats = 0
	w.WriteHeader(http.StatusOK)
func (c *clientCoordinator) pathGetShards(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		c.log.Info("GET 401 api/shards")
		return
	}

	bots := c.Bots
	for _, bot := range bots {
		bot.Token = ""
	}

	jsons, err := json.Marshal(bots)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Errorf("GET 500 api/shards: %s", err)
		return
	}
	w.Write(jsons)
	c.log.Info("GET 200 api/shards")
}
