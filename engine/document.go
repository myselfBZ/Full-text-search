package engine

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocument(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	decoder := xml.NewDecoder(gz)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := decoder.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil

}
