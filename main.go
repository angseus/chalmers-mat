package main

import (
    "log"
    "net/http"  
)

func main() {
    // another goroutine to update restaurants
    go UpdateRestaurants()

    // create router
    router := NewRouter()

    // start serving
    log.Fatal(http.ListenAndServe(":8080", router))
}
