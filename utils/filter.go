package utils

import(
    "strings"
    snowballeng "github.com/kljensen/snowball/english"
)



func lowercaseFilter(text []string) []string {
    r := make([]string, len(text))
    for i, token := range text{
        r[i] = strings.ToLower(token)
    }
    return r 
}

func stopwordFilter(tokens []string) []string  {
    stopwords := map[string]struct{}{
        "a":{},"and":{},"be":{},"have":{},"i":{},
        "in":{},"of":{},"that":{},"the":{}, "to":{},
    }
    r := make([]string, len(tokens))
    for _, token := range tokens{
        if _, ok := stopwords[token]; !ok{
            r = append(r, token)
        }
    }
    return r 
}

func stemmerFilter(tokens []string) []string{
    r := make([]string, len(tokens))
    for i, token := range tokens{
        r[i] = snowballeng.Stem(token, false)
    }
    return r 
}





