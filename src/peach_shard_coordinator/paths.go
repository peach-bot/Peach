package main

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Function used by clients to get acquire a shard number for login
func getShard(w http.ResponseWriter, r *http.Request) {
	// response setup
	w.Header().Set("Content-Type", "application/json")

	// get next unreserved shard
	for pos, thisshard := range shards {
		if thisshard.Reserved == false {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`{"total_shards": %d, "assigned_shard": %d, "gatewayurl": "%s"}`, len(shards), thisshard.ShardID, gatewayurl)))
			log.WithFields(log.Fields{
				"total_shards": len(shards),
				"shardID":      thisshard.ShardID,
				"active":       shards[thisshard.ShardID].Active,
			}).Info("GET 200 api/v1/getshard - shard assigned")
			break
		}
		if pos == (len(shards) - 1) {
			w.WriteHeader(http.StatusNoContent)
			log.Info("GET 204 api/v1/getshard - all shards assigned")
		}
	}
}

// Function used to reserve a shard number for a client
func reserveShard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shardID, err := strconv.Atoi(r.URL.Query().Get("shardid"))
	if err != nil {
		log.Fatal(err)
	}
	/*
		Because of the split shard 0 shard 1 has the index 2 in the shards list.
		When using the api -1 refers to the DM shard and 0 refers to the server shard 0. 1 then refers to the server shard 1.
	*/

	if shardID >= len(shards) || shardID < 0 {
		// requested shard out of range
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Requested shard out of range"))
		log.WithFields(log.Fields{
			"shardID": shardID,
		}).Info("POST 406 api/v1/reserveshard - requested shard out of range")
	} else if shards[shardID].Reserved {
		// shard already reserved
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Requested shard already reserved"))
		log.WithFields(log.Fields{
			"shardID":  shardID,
			"reserved": shards[shardID].Reserved,
			"active":   shards[shardID].Active,
		}).Info("POST 406 api/v1/reserveshard - requested shard already reserved")
	} else {
		// set shard reservation
		shards[shardID].Reserved = true
		w.WriteHeader(http.StatusCreated)
		log.WithFields(log.Fields{
			"shardID":   shardID,
			"reserved":  shards[shardID].Reserved,
			"active":    shards[shardID].Active,
			"is_server": shards[shardID].Server,
		}).Info("POST 201 api/v1/reserveshard - shard reserved")
	}
}

// Function used to update the state of a single shard
func updateShard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
}

// Function to reset shards object and update shard amount
func scale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		log.Fatal(err)
	}
	go resetShardCount(amount)
}
