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
    results := engine.Idx.Search(searchRequest.Text)
    response := SearchResponse{
        Results: make([]engine.Document, 0, len(results)),
    }
    for _, id := range results{
        response.Results = append(response.Results, engine.Docs[id]) 
    }
    log.Println(response)
}

