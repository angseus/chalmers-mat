package main

import (
	"github.com/mmcdole/gofeed"
	"time"
	"strings"
)

type ExpressMenu struct {
	Express 	string	`json:"express"`
	Veg 		string 	`json:"veg"`
}

// let it be global for now
var express ExpressMenu = ExpressMenu{}

func FetchExpress() ExpressMenu{
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://intern.chalmerskonferens.se/view/restaurant/express/V%C3%A4nster.rss")
	today := time.Now().String()[0:10]

	for _, day := range feed.Items {
		if (strings.Split(day.Title, "Meny Express - ")[1]) == today{
			express.Express = strings.Split(strings.Split(day.Description, "<td>")[1], "</td>")[0]
			express.Veg = strings.Split(strings.Split(day.Description, "<td>")[2], "</td>")[0]
		}
	}

	return express

}