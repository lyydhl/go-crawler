package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
)

const profileRe = `<div class="m-btn purple"[^>]*>([^<]+)</div>`

func ParserProfile(contents []byte, name string) engine.ParserResult {
	re := regexp.MustCompile(profileRe)
	matchs := re.FindAllSubmatch(contents, -1)

	profile := model.Profile{}
	profile.Name = name
	for i, m := range matchs {
		if i == 0 {
			profile.Marriage = string(m[1])
		} else if i == 1 {
			profile.Age = string(m[1])
		} else if i == 2 {
			profile.Constellation = string(m[1])
		} else if i == 3 {
			profile.Height = string(m[1])
		} else if i == 4 {
			profile.Weight = string(m[1])
		} else if i == 5 {
			profile.Workplace = string(m[1])
		} else if i == 6 {
			profile.Income = string(m[1])
		} else if i == 7 {
			profile.Occupation = string(m[1])
		} else if i == 8 {
			profile.Education = string(m[1])
		}
	}

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}
