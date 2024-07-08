package main

import (
	"log"
	"net/http"

	"github.com/myselfBZ/full-text-search/web"
)




func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/search", web.HandleRequest)
    log.Fatal(http.ListenAndServe(":8080", mux))
}
