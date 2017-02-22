package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// variables for status
var calls int64 = 0
var startTime = time.Now()

// middleware for API calls
func APIMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// keep some statistics
		calls = calls + 1

		// proceed to the actual handler
		h.ServeHTTP(w, r)
	})
}

func APIStatus(w http.ResponseWriter, r *http.Request) {

	s := StatusMessage{"Up and running!", time.Since(startTime).String(), calls}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}