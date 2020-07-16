package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

/*
// TestGetShard ensures the /getshard path works properly.
func TestGetShard(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount, true)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Run request
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))

	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 0, "gatewayurl": ""}`, wantedShardCount), w.Body.String())
}

// TestReserveShard ensures the /reserveshard path works properly.
func TestReserveShard(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount, true)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)

	// Run request
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))

	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)
}

// TestGetAndReserveMultipleShards ensures that shards get assigned and reserved correctly.
func TestGetAndReserveMultipleShards(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 3
	resetShardCount(wantedShardCount, true)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Get Shard 0
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 0, "gatewayurl": ""}`, wantedShardCount), w.Body.String())

	// Reserve Shard 0
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Get Shard 1
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 1, "gatewayurl": ""}`, wantedShardCount), w.Body.String())

	// Reserve Shard 1
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=1", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Get Shard 2
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(fmt.Sprintf(`{"total_shards": %v, "assigned_shard": 2, "gatewayurl": ""}`, wantedShardCount), w.Body.String())

	// Reserve Shard 2
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=2", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)
}

// TestReserveOutOfBounds ensures that nonexistant shards can not get reserved
func TestReserveOutOfBounds(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount, true)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Reserve Shard 2 (too high)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=2", nil))
	// Evaluate results
	assert.Equal(http.StatusNotAcceptable, w.Code)

	// Reserve Shard -1 (too low)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=-1", nil))
	// Evaluate results
	assert.Equal(http.StatusNotAcceptable, w.Code)
}

func TestAllShardsAssigned(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount, true)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Reserve Shard 0
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Reserve Shard 1
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=1", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Get Shard (All assigned)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/getshard", nil))
	// Evaluate results
	assert.Equal(http.StatusNoContent, w.Code)
}

func TestShardAlreadyAssigned(t *testing.T) {
	// Test setup
	assert := assert.New(t)
	wantedShardCount := 2
	resetShardCount(wantedShardCount, true)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)

	// Reserve Shard 0
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))
	// Evaluate results
	assert.Equal(http.StatusCreated, w.Code)

	// Reserve Shard 0
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/api/v1/reserveshard?shardid=0", nil))
	// Evaluate results
	assert.Equal(http.StatusNotAcceptable, w.Code)
}
*/
