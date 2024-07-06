package main

import (
    "net/http"
    "github.com/myselfBZ/full-text-search/utils"

   "time"
)

type Stats struct{
    LoadDocumentTime time.Time `json:"documentLoadTime"`
    IndexDocumentNumber int `json:"indexDocumentTime"`

}

type SearchRequest struct{
    Text string `json:"text"`
}

func main() {

}
