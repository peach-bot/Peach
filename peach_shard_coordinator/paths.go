package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Function used by clients to get acquire a shard number for login
func getShard(w http.ResponseWriter, r *http.Request) {
	// response setup
	w.Header().Set("Content-Type", "application/json")

	// get next unreserved shard
	for pos, thisshard := range shards {
		if thisshard.Reserved == false {
			// write response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`{"total_shards": %d, "assigned_shard": %d, "is_server": %v}`, len(shards), thisshard.ShardID, thisshard.Server)))
			break
		}
		if pos == len(shards) {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}

// Function used to reserve a shard number for a client
func reserveShard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)
	shardID, err := strconv.Atoi(pathParams["shardID"])
	if err != nil {
		log.Fatal(err)
	}
	/*
		Because of the split shard 0 shard 1 has the index 2 in the shards list.
		When using the api -1 refers to the DM shard and 0 refers to the server shard 0. 1 then refers to the server shard 1.
	*/
	shardID++

	// set shard reservation
	shards[shardID].Reserved = true
	w.WriteHeader(http.StatusCreated)
}

// Function used to update the state of a single shard
func updateShard(w http.ResponseWriter, r *http.Request) {

}

// Function to reset shards object and update shard amount
func scale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	go resetShardCount()
}
