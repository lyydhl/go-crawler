package engine

import (
	"fmt"
	"go-crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seed ...Request) {
	var requests []Request

	// for _, r := range seed {
	// 	requests = append(requests, r)
	// }
	requests = append(requests, seed...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			fmt.Printf("Got item %v\n", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error Fecth Url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFun(body), nil
}
