package main

/***********/
/* Imports */
/***********/
import (
    "encoding/json"
    "net/http"
    "time"
)

/*********/
/* Types */
/*********/
// struct for errors sent using JSON
type jsonErr struct {
    Code int    `json:"code"`
    Text string `json:"text"`
}

// struct for status JSON
type StatusMessage struct {
    Status string   `json:"status"`
    Uptime string   `json:"uptime"`
    Calls int64     `json:"calls"`
}

/*************/
/* Variables */
/*************/
// api routes
var routes = Routes{
    Route{
        "APIStatus",
        "GET",
        "/api",
        APIStatus,
    },
    Route{
        "Express",
        "GET",
        "/api/express",
        Express,
    },
    Route{
        "Kårrestaurangen",
        "GET",
        "/api/karis",
        Karis,
    },
    Route{
        "Linsen",
        "GET",
        "/api/linsen",
        Linsen,
    },
    Route{
        "J.A. Pripps",
        "GET",
        "/api/karis",
        JA,
    },
    Route{
        "Hyllan",
        "GET",
        "/api/hyllan",
        Hyllan,
    },
    Route{
        "Kokboken",
        "GET",
        "/api/kokboken",
        Kokboken,
    },
    Route{
        "L's Resto",
        "GET",
        "/api/lsresto",
        Lsresto,
    },
    Route{
        "L's Kitchen",
        "GET",
        "/api/lskitchen",
        Lskitchen,
    },
    Route{
        "Einstein",
        "GET",
        "/api/einstein",
        Einstein,
    },
}

// variables for status
var calls int64 = 0
var startTime = time.Now()

/*****************/
/* API FUNCTIONS */
/*****************/
// middleware for API calls
func APIMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // keep some statistics
        calls = calls + 1

        // proceed to the actual handler
        h.ServeHTTP(w, r)
    })
}

// API Status
func APIStatus(w http.ResponseWriter, r *http.Request) {
    s := StatusMessage{"Up and running!", time.Since(startTime).String(), calls}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(s); err != nil {
        panic(err)
    }
}

// Express
func Express(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(express); err != nil {
        panic(err)
    }
}

// Kårrestaurangen
func Karis(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(karis); err != nil {
        panic(err)
    }
}

// Linsen
func Linsen(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(linsen); err != nil {
        panic(err)
    }
}

// J.A. Pripps
func JA(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(ja); err != nil {
        panic(err)
    }
}

// Hyllan
func Hyllan(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(hyllan); err != nil {
        panic(err)
    }
}

// Kokboken
func Kokboken(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(kokboken); err != nil {
        panic(err)
    }
}

// L's Resto
func Lsresto(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(lsresto); err != nil {
        panic(err)
    }
}

// L's Kitchen
func Lskitchen(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(lskitchen); err != nil {
        panic(err)
    }
}

// Einstein
func Einstein(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(einstein); err != nil {
        panic(err)
    }
}
