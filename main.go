package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	
	// test to fetch express menu
	fmt.Println("Fetching Express menu!")
	FetchExpress()
	fmt.Println(express.Express)
	fmt.Println(express.Veg)

	// create router
	router := NewRouter()

	// start serving
	log.Fatal(http.ListenAndServe(":8080", router))
}
