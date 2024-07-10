package main

import (
	"log"
	"net/http"

	"github.com/myselfBZ/full-text-search/engine"
	"github.com/myselfBZ/full-text-search/web"
)



func main() {
    engine.Initialze()
    mux := http.NewServeMux()
    mux.HandleFunc("/search", web.HandleRequest)
    log.Println("Starting the server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
