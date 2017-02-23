package main

// struct for errors sent using JSON
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// struct for status JSON
type StatusMessage struct {
    Status string	`json:"status"`
    Uptime string	`json:"uptime"`
    Calls int64		`json:"calls"`
}

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
}