package engine

import "log"

var Idx = &Index{}

var Docs = []Document{}

func Initialze() {
	Docs, err := LoadDocument("../enwiki-latest-abstract.xml.gz")
	if err != nil {
		log.Fatal(err)
	}
    Idx.Add(Docs)
}
