package engine

import (
	"fmt"
	"go-crawler/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request

	// for _, r := range seed {
	// 	requests = append(requests, r)
	// }
	requests = append(requests, seed...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch error Fecth Url %s: %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFun(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			fmt.Printf("Got item %v\n", item)
		}
	}
}
