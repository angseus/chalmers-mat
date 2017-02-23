package main

import (
	"time"
	"strings"
	"github.com/mmcdole/gofeed"
	"github.com/PuerkitoBio/goquery"
)

// struct for restaurant
type Restaurant struct {
	Name 		string 	`json:"name"` // name of the restaurant
	Entries 	[]Entry `json:"entries"` // list of dishes for that restaurant
}

// struct for entry in restaurant
type Entry struct {
	Name	string	`json:"name"` // name of the entry
	Dish	string  `json:"dish"` // name of the dish
}

// globals for all restaurants
var karis Restaurant = Restaurant{}
var express Restaurant = Restaurant{}
var linsen Restaurant = Restaurant{}
var ja Restaurant = Restaurant{}
var hyllan Restaurant = Restaurant{}
var kokboken Restaurant = Restaurant{}
var lsresto Restaurant = Restaurant{}
var lskitchen Restaurant = Restaurant{}
var einstein Restaurant = Restaurant{}


func UpdateRestaurants(){
	for {
    	karis = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss")
		express = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/express/V%C3%A4nster.rss")
		linsen = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/linsen/RSS%20Feed.rss")
		ja = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/j-a-pripps-pub-cafe/RSS%20Feed.rss")
		hyllan = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/hyllan/RSS%20Feed.rss")
		kokboken = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/kokboken/RSS%20Feed.rss")
		lsresto = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/l-s-resto/RSS%20Feed.rss")
		lskitchen = FetchCHS("http://intern.chalmerskonferens.se/view/restaurant/l-s-kitchen/Projektor.rss")
		einstein = FetchEinstein()
		RestaurantLogger("Updated all restaurants! Waiting 15 minutes until next update")
		time.Sleep(15 * time.Minute)
	}
}


func FetchCHS(rss string) Restaurant {
	// get the rss feed
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)
	
	restaurantName := (strings.TrimRight(strings.Split(feed.Title, "Meny ")[1], " "))

	// create Restaurant structure
	restaurant := Restaurant{restaurantName, nil}

	// get todays date
	today := time.Now().String()[0:10]

	for _, day := range feed.Items {
		// check if current date
		if (strings.Split(day.Title, "Meny " + restaurantName + " - ")[1]) == today{
			// go through all lines in rss
			for _, line := range strings.Split(day.Description, "\n") {
				// if name of entry
   				if(strings.Contains(line, "<b>")){
   					restaurant.Entries = append(restaurant.Entries, Entry{strings.TrimSpace(strings.Split(strings.Split(line, "<b>")[1], "</b>")[0]), ""})
   				}
   				// if name of dish
   				// we know that it should be added to the previously added entry
   				if(strings.Contains(line, "<td>")){
   					restaurant.Entries[len(restaurant.Entries)-1].Dish = strings.TrimSpace(strings.Split(strings.Split(line, "<td>")[1], "</td>")[0])
   				}
			}
		}
	}
	RestaurantLogger("Updated: " + restaurantName)
	return restaurant
}

func FetchEinstein() Restaurant {
	einstein := Restaurant{"Einstein", nil}

	// fetch the document
	doc, err := goquery.NewDocument("http://www.butlercatering.se/einstein")
    if err != nil {
        panic(err)
    }

    // translate days to Swedish
    day := time.Now().Local().Weekday().String()
    switch day {
	case "Monday":
		day = "Måndag"
	case "Tuesday":
		day = "Tisdag"
	case "Wednesday":
		day = "Onsdag"
    case "Thursday":
    	day = "Torsdag"
	case "Friday":
		day = "Fredag"
    }
    
    // find proper tag
    doc.Find(".field-day").Each(func(i int, s *goquery.Selection) {
    	// check if it is the current day
    	if (strings.Contains(s.Text(), day)){
    		counter := 0
    		for _, line := range strings.Split(s.Text(), "\n") {
    			tmp := strings.TrimSpace(line)

    			if (counter == 2){
					einstein.Entries = append(einstein.Entries, Entry{"Lunch", tmp})
				}
				if (counter == 3){
					einstein.Entries = append(einstein.Entries, Entry{"Lunch", tmp})
				}
				if (counter == 4){
					tmp = strings.TrimSpace(strings.Split(tmp, "Veg:")[1])
					einstein.Entries = append(einstein.Entries, Entry{"Vegetariskt", tmp})
				}
				if (counter == 5){
					einstein.Entries = append(einstein.Entries, Entry{"Special", tmp})
				}
				counter++
    		}
    	}
  	})
  	RestaurantLogger("Updated: Einstein")
	return einstein
}