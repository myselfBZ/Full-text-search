package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/myselfBZ/full-text-search/engine"
)

func HandleRequest(w http.ResponseWriter, r *http.Request)  {
    var searchRequest SearchRequest
    if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    log.Print("Decoded successfully\n")
    results := engine.Idx.Search(searchRequest.Text)
    log.Println("searching...")
    response := SearchResponse{
        Results: make([]engine.Document, 80, len(results)),
    }
    log.Println("Allocated mem for response")
    for _, id := range results{
        log.Println("Appending results...")
        response.Results = append(response.Results, engine.Docs[id]) 
    }
    json.NewEncoder(w).Encode(response)
}

