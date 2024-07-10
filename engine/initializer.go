package engine

import (
	"log"
	"time"
)

var idx = make(Index)
var Idx = &idx
var Docs = []Document{}

func Initialze() {
    start := time.Now()
	Docs, err := LoadDocument("enwiki-latest-abstract.xml.gz")
	if err != nil {
		log.Fatal(err)
	}
    log.Println("Loading the documents took: ", time.Since(start))
    start = time.Now()
    Idx.Add(Docs)
    log.Println("Indexing documents took: ", time.Since(start))
    
}
