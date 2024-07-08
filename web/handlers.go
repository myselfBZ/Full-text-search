package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/myselfBZ/full-text-search/engine"
)

func HandleRequest(w http.ResponseWriter, r *http.Request)  {
    var searchRequest SearchRequest
    if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    start := time.Now()
    results := engine.Idx.Search(searchRequest.Text)
    response := SearchResponse{}
    for _, id := range results{
        response.Results = append(response.Results, engine.Docs[id]) 
    }
    stop := time.Since(start)
    response.Time = stop 
    json.NewEncoder(w).Encode(response)
}

