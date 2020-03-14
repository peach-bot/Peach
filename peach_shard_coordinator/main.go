package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var shards []shard
var bottoken string = os.Getenv("BOTTOKEN")

// function used to create the shards object and to fetch the shard amount
func resetShardCount() {

	// fetch recommended shard amount from discord api
	client := &http.Client{}
	var response gatewayresponse
	req, err := http.NewRequest("GET", "https://discordapp.com/api/gateway/bot", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bot %v", bottoken))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&response)

	// create list with shard objects
	shardCount := response.Shards + 1
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

func main() {
	log.Printf("shard coordinator online\n")

	// setup paths
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)
	api.HandleFunc("/reserveshard/{shardID}", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/updateshard", updateShard).Methods(http.MethodPost)
	api.HandleFunc("/scale", scale).Methods(http.MethodGet)

	// initial creation
	go resetShardCount()

	// run
	log.Fatal(http.ListenAndServe(":8080", r))
}
