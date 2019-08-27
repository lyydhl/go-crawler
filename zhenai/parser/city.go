package parser

import (
	"go-crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

const avatarRe = `<div class="photo"><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank"><img src="([^?]+)?[^"]*" alt="([^"]+)"></a></div>`

func ParserCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(avatarRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range matchs {
		avatar := string(m[2])
		name := string(m[3])
		result.Items = append(result.Items, name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFun: func(c []byte) engine.ParserResult {
					return ParserProfile(c, name, avatar)
				},
			})
	}
	return result
}
