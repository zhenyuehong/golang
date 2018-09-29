package engine

import (
	"golang/carwler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	//print url
	//log.Printf("fetching url: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher error: fether url %s, %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
	//return r.Parser.Parse(body, r.Url), nil
}
