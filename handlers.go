package main

import (
    "net/http"
    "encoding/json"
)

func HandleRequest(w http.ResponseWriter, r *http.Request)  {
    var searchRequest SearchRequest
    if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    w.WriteHeader(http.StatusOK)
}
