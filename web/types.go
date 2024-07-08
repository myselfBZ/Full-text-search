package web

import (
	"time"

	"github.com/myselfBZ/full-text-search/engine"
)


type SearchRequest struct{
    Text string `json:"text"`
}

type SearchResponse struct{
    Results []engine.Document `json:"result"`
    Time time.Duration `json:"time"`
}
