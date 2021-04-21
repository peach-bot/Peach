package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

	w.WriteHeader(http.StatusUnauthorized)

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
	if bot == nil {
		return nil, nil, errors.New("invalid bot_id")
	}

	shard := bot.Shards[shardID]
	if shard == nil {
		return nil, nil, errors.New("invalid shard_id")
	}

	w.WriteHeader(http.StatusOK)
	return bot, shard, nil
}

func (c *clientCoordinator) pathLogin(w http.ResponseWriter, r *http.Request) {
	c.log.Debug("GET called api/login")
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Infof("GET 401 api/login: %s", err)
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	bot, shard := c.nextShard()
	if bot == nil || shard == nil {
		w.WriteHeader(http.StatusNoContent)
		c.log.Info("GET 204 api/ready: all shards assigned")
		return
	}

	go c.shardManager(bot, shard)

	response := fmt.Sprintf(`{"token": "%s", "total_shards": %d, "assigned_shard": %d, "gateway_url": "%s", "heartbeat_interval": "%s"}`, bot.Token, bot.ShardCount, shard.ShardID, c.GatewayURL, c.heartbeatInterval)
	shard.Reserved = true
	shard.LastHeartbeat = time.Now()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	c.log.Info("GET 200 api/login")
}

func (c *clientCoordinator) pathReady(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Info("GET 401 api/ready")
		return
	}
	_, shard, err := c.verifyBotShard(w, r)
	if err != nil {
		w.Write([]byte(err.Error()))
		c.log.Info("GET 404 api/ready")
		return
	}
	if shard.Active {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Shard already active"))
		c.log.Info("GET 403 api/ready")
		return
	}

	shard.Active = true
	w.WriteHeader(http.StatusOK)
	c.log.Info("GET 200 api/ready")
}

func (c *clientCoordinator) pathHeartbeat(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Info("GET 401 api/heartbeat")
		return
	}
	_, shard, err := c.verifyBotShard(w, r)
	if err != nil {
		w.Write([]byte(err.Error()))
		c.log.Info("GET 404 api/heartbeat")
		return
	}
	shard.LastHeartbeat = time.Now()
	shard.MissedHeartbeats = 0
	w.WriteHeader(http.StatusOK)
	c.log.Info("GET 200 api/heartbeat")
}

func (c *clientCoordinator) pathGetGuildSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Infof("GET 401 api/guilds/%s", vars["guildID"])
		return
	}
	_, _, err = c.verifyBotShard(w, r)
	if err != nil {
		w.Write([]byte(err.Error()))
		c.log.Infof("GET 404 api/guilds/%s", vars["guildID"])
		return
	}

	s, err := db.getGuildSettings(vars["guildID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Errorf("GET 500 api/guilds/%s: %s", vars["guildID"], err)
		return
	}

	jsons, err := json.Marshal(*s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Errorf("GET 500 api/guilds/%s: %s", vars["guildID"], err)
		return
	}
	w.Write(jsons)
	c.log.Infof("GET 200 api/guilds/%s", vars["guildID"])
}

func (c *clientCoordinator) pathGetUserSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Infof("GET 401 api/users/%s", vars["userID"])
		return
	}
	_, _, err = c.verifyBotShard(w, r)
	if err != nil {
		w.Write([]byte(err.Error()))
		c.log.Infof("GET 404 api/users/%s", vars["userID"])
		return
	}

	s, err := db.getUserSettings(vars["userID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Errorf("GET 500 api/users/%s: %s", vars["userID"], err)
		return
	}

	jsons, err := json.Marshal(*s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Errorf("GET 500 api/users/%s: %s", vars["userID"], err)
		return
	}
	w.Write(jsons)
	c.log.Infof("GET 200 api/users/%s", vars["userID"])
}

func (c *clientCoordinator) pathGetShards(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		c.log.Info("GET 401 api/shards")
		return
	}

	bots := c.Bots
	// for _, bot := range bots {
	// 	bot.Token = ""
	// }

	jsons, err := json.Marshal(bots)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c.log.Errorf("GET 500 api/shards: %s", err)
		return
	}
	w.Write(jsons)
	c.log.Info("GET 200 api/shards")
}
