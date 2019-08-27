package parser

import (
	"go-crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	limitCount := 0

	for _, m := range matchs {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m[1]),
				ParserFun: ParserCity,
			})
		limitCount++
		if limitCount > 3 {
			break
		}
		//fmt.Printf("city: %s  url: %s\n", m[2], m[1])
	}
	//fmt.Println("matchs count: ", len(matchs))
	return result
}
