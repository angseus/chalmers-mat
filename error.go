package main

// struct for errors sent using JSON
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
