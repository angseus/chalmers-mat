package main

// struct for status JSON
type StatusMessage struct {
    Status string	`json:"status"`
    Uptime string	`json:"uptime"`
    Calls int64		`json:"calls"`
}