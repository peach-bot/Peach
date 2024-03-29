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

func (c *Coordinator) verifyAuth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	authsecret := r.Header.Get("authorization")
	if authsecret != c.Config.Secret {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New(("Auth code didn't match cluster secret"))
	}
	return nil
}

func (c *Coordinator) verifyBotShard(w http.ResponseWriter, r *http.Request) (*Bot, *Shard, error) {

	botID := r.Header.Get("bot_id")
	if botID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, nil, errors.New("Header missing bot_id")
	}

	sid := r.Header.Get("shard_id")
	if sid == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, nil, errors.New("Header missing shard_id")
	}

	shardID, err := strconv.Atoi(sid)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, nil, errors.New("invalid shard_id")
	}

	bot := c.Bots[botID]
	if bot == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, nil, errors.New("invalid bot_id")
	}

	shard := bot.Shards[shardID]
	if shard == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, nil, errors.New("invalid shard_id")
	}

	return bot, shard, nil
}

func (c *Coordinator) pathLogin(w http.ResponseWriter, r *http.Request) {
	c.log.Debug("GET called api/login")

	// verify Authentication
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Infof("GET 401 api/login: %s", err)
		return
	}

	switch r.Header.Get("type") {
	case "client":
		// Allocate shard
		c.lock.Lock()
		defer c.lock.Unlock()

		bot, shard := c.nextShard()
		if bot == nil || shard == nil {
			w.WriteHeader(http.StatusNoContent)
			c.log.Info("GET 204 api/login: all shards assigned")
			return
		}
		go c.shardManager(bot, shard)
		shard.Reserved = true
		shard.LastHeartbeat = time.Now()

		// send response to client
		response := fmt.Sprintf(`{"token": "%s", "total_shards": %d, "assigned_shard": %d, "gateway_url": "%s", "heartbeat_interval": "%s", "spotify_client_id": "%s", "spotify_client_secret": "%s"}`, bot.Token, bot.ShardCount, shard.ShardID, c.GatewayURL, c.heartbeatInterval, c.Config.Coordinator.SpotifyClientID, c.Config.Coordinator.SpotifyClientSecret)
		w.Write([]byte(response))

	case "launcher":

		c.lock.Lock()
		defer c.lock.Unlock()

		id := r.Header.Get("id")

		l := new(Launcher)
		l.ActiveClients = 0
		l.ID = id
		maxc := r.Header.Get("max_clients")
		maxClients, err := strconv.Atoi(maxc)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid max_clients"))
		}
		l.MaxClients = maxClients
		c.Launchers[id] = l

		// send response to launcher
		response := fmt.Sprintf(`{"heartbeat_interval": "%s"}`, c.heartbeatInterval)
		w.Write([]byte(response))

	}

	w.WriteHeader(http.StatusOK)
	c.log.Info("GET 200 api/login")

}

func (c *Coordinator) pathReady(w http.ResponseWriter, r *http.Request) {
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

func (c *Coordinator) pathHeartbeat(w http.ResponseWriter, r *http.Request) {
	err := c.verifyAuth(w, r)
	if err != nil {
		c.log.Info("GET 401 api/heartbeat")
		return
	}

	switch r.Header.Get("type") {
	case "client":
		_, shard, err := c.verifyBotShard(w, r)
		if err != nil {
			w.Write([]byte(err.Error()))
			c.log.Info("GET 404 api/heartbeat")
			return
		}
		shard.LastHeartbeat = time.Now()
		shard.MissedHeartbeats = 0
	}

	w.WriteHeader(http.StatusOK)
	c.log.Info("GET 200 api/heartbeat")
}

func (c *Coordinator) pathGetGuildSettings(w http.ResponseWriter, r *http.Request) {
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

	s, err := c.DB.getGuildSettings(vars["guildID"])
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

func (c *Coordinator) pathGetUserSettings(w http.ResponseWriter, r *http.Request) {
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

	s, err := c.DB.getUserSettings(vars["userID"])
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

func (c *Coordinator) pathGetShards(w http.ResponseWriter, r *http.Request) {
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
