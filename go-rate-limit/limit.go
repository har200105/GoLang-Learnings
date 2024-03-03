package main

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func tokenBucketRateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(5, 7)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body:   "API Rate Limitted",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}

		next(w, r)

	})
}

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mutex   sync.Mutex
	clients = make(map[string]*client)
)

func perClientRateLimitter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {

	go func() {
		for {
			time.Sleep(time.Minute)
			mutex.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mutex.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mutex.Lock()
		if _, found := clients[ipAddress]; !found {
			clients[ipAddress] = &client{limiter: rate.NewLimiter(5, 7)}
		}
		clients[ipAddress].lastSeen = time.Now()
		if !clients[ipAddress].limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body:   "API Rate Limitted",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}
		mutex.Unlock()
		next(w, r)
	})
}
