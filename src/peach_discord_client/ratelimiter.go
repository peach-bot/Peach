package main

import (
	"math"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Ratelimiter struct {
	sync.Mutex
	Routes map[string]*Route
	global *int64
}

type Route struct {
	sync.Mutex
	Route     string
	Limit     int
	Remaining int
	Reset     time.Time
	global    *int64
}

func CreateRatelimiter() *Ratelimiter {
	return &Ratelimiter{
		Routes: make(map[string]*Route),
		global: new(int64),
	}
}

func (r *Ratelimiter) GetRoute(routeid string) *Route {
	r.Lock()
	defer r.Unlock()

	if route, ok := r.Routes[routeid]; ok {
		return route
	}

	route := &Route{Remaining: 1, Route: routeid, global: r.global}
	r.Routes[routeid] = route
	return route
}

func (r *Route) WaitTime() time.Duration {

	if r.Remaining < 1 && r.Reset.After(time.Now()) {
		return r.Reset.Sub(time.Now())
	}

	sleeptill := time.Unix(0, atomic.LoadInt64(r.global))
	if time.Now().Before(sleeptill) {
		return sleeptill.Sub(time.Now())
	}

	return 0
}

func (r *Route) Wait() {

	w := r.WaitTime()
	if w > 0 {
		time.Sleep(w)
	}

}

func (r *Route) Prepare() {
	r.Lock()
	r.Wait()
	r.Remaining--
}

func (r *Ratelimiter) PrepareRoute(routeid string) *Route {
	route := r.GetRoute(routeid)
	route.Prepare()
	return route
}

func (r *Route) Release(h http.Header) error {
	defer r.Unlock()

	if h == nil {
		return nil
	}

	remaining := h.Get("X-RateLimit-Remaining")
	reset := h.Get("X-RateLimit-Reset")
	global := h.Get("X-RateLimit-Global")
	resetAfter := h.Get("X-RateLimit-Reset-After")

	if resetAfter != "" {
		parsedAfter, err := strconv.ParseFloat(resetAfter, 64)
		if err != nil {
			return err
		}

		whole, frac := math.Modf(parsedAfter)
		resetAt := time.Now().Add(time.Duration(whole) * time.Second).Add(time.Duration(frac*1000) * time.Millisecond)

		if global != "" {
			atomic.StoreInt64(r.global, resetAt.UnixNano())
		} else {
			r.Reset = resetAt
		}
	} else if reset != "" {
		discordTime, err := http.ParseTime(h.Get("Date"))
		if err != nil {
			return err
		}

		unix, err := strconv.ParseFloat(reset, 64)
		if err != nil {
			return err
		}

		whole, frac := math.Modf(unix)
		delta := time.Unix(int64(whole), 0).Add(time.Duration(frac*1000)*time.Millisecond).Sub(discordTime) + time.Millisecond*250
		r.Reset = time.Now().Add(delta)
	}

	if remaining != "" {
		parsedRemaining, err := strconv.ParseInt(remaining, 10, 32)
		if err != nil {
			return err
		}
		r.Remaining = int(parsedRemaining)
	}

	return nil
}
