package main

import (
	"log"
	"net/http"
	"fmt"
	
)

var karis Restaurant = Restaurant{}
var express Restaurant = Restaurant{}
var linsen Restaurant = Restaurant{}
var ja Restaurant = Restaurant{}
var hyllan Restaurant = Restaurant{}
var kokboken Restaurant = Restaurant{}
var lsresto Restaurant = Restaurant{}
var lskitchen Restaurant = Restaurant{}
var einstein Restaurant = Restaurant{}

func main() {
	// karis
	karis = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss")
	fmt.Println(karis.Entries)
	fmt.Println("")

	// express
	express = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/express/V%C3%A4nster.rss")
	fmt.Println(express.Entries)
	fmt.Println("")

	linsen = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/linsen/RSS%20Feed.rss")
	fmt.Println(linsen.Entries)
	fmt.Println("")

	ja = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/j-a-pripps-pub-cafe/RSS%20Feed.rss")
	fmt.Println(ja.Entries)
	fmt.Println("")

	hyllan = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/hyllan/RSS%20Feed.rss")
	fmt.Println(hyllan.Entries)
	fmt.Println("")

	kokboken = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/kokboken/RSS%20Feed.rss")
	fmt.Println(kokboken.Entries)
	fmt.Println("")

	lsresto = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/l-s-resto/RSS%20Feed.rss")
	fmt.Println(lsresto.Entries)
	fmt.Println("")

	lskitchen = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/l-s-kitchen/Projektor.rss")
	fmt.Println(lskitchen.Entries)
	fmt.Println("")
	
	einstein = FetchEinstein()
	fmt.Println(einstein.Entries)
	fmt.Println("")

	// create router
	router := NewRouter()

	// start serving
	log.Fatal(http.ListenAndServe(":8080", router))
}
