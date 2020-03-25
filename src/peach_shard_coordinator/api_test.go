package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetShard(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Run request
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))

	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 0, "api_shardid": -1, "is_server": false}`, wantedShardCount), w.Body.String())
}

func TestReserveShard(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)

	// Run request
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=-1", nil))

	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)
}

func TestMultiGetReserveShard(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Get Shard 0 (DM)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 0, "api_shardid": -1, "is_server": false}`, wantedShardCount), w.Body.String())

	// Reserve Shard 1
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=1", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Get Shard 0 (DM)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 0, "api_shardid": -1, "is_server": false}`, wantedShardCount), w.Body.String())

	// Reserve Shard 0 (DM)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=-1", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Reserve Shard 10 (Out of bounds)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=10", nil))
	// Evaluate results
	assert.Equal(http.StatusNotAcceptable, w.Code)

	// Get Shard 0 (Server)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 0, "api_shardid": 0, "is_server": true}`, wantedShardCount), w.Body.String())

	// Reserve Shard 0 (Server)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Reserve Shard 0 (Server, already assigned)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))
	// Evaluate results
	assert.Equal(http.StatusNotAcceptable, w.Code)

	// Get Shard (All assigned)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusNoContent, w.Code)
}
