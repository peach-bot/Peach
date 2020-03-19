package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// function used to create the shards object and to fetch the shard amount
func resetShardCount(shardCount int) {
	if shardCount == 0 {
		// fetch recommended shard amount from discord api
		client := &http.Client{}
		var response gatewayresponse
		req, err := http.NewRequest("GET", "https://discordapp.com/api/gateway/bot", nil)
		if err != nil {
			log.Error(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bot %v", bottoken))
		resp, err := client.Do(req)
		if err != nil {
			log.Error(err)
		}
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&response)

		shardCount = response.Shards + 1
	} else {
		shardCount++
	}

	// create list with shard objects
	shards = make([]shard, (shardCount))

	// set shardIDs and roles
	DMshard := true
	for shardID := 0; shardID < shardCount; shardID++ {
		if DMshard {
			shards[shardID].ShardID = shardID
			DMshard = false
		} else {
			shards[shardID].ShardID = shardID - 1
			shards[shardID].Server = true
		}
	}
}
