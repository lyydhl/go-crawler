package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"go-crawler/util"
	"regexp"
	"strings"
)

// const profileRe = `<div class="m-btn purple"[^>]*>([^<]+)</div>`

// func ParserProfile(contents []byte, name string) engine.ParserResult {
// 	re := regexp.MustCompile(profileRe)
// 	matchs := re.FindAllSubmatch(contents, -1)

// 	profile := model.Profile{}
// 	profile.Name = name
// 	for i, m := range matchs {
// 		if i == 0 {
// 			profile.Marriage = string(m[1])
// 		} else if i == 1 {
// 			profile.Age = util.StringToIntDefault(strings.Replace(string(m[1]), "岁", "", -1))
// 		} else if i == 2 {
// 			profile.Constellation = string(m[1])
// 		} else if i == 3 {
// 			profile.Height = util.StringToIntDefault(strings.Replace(string(m[1]), "cm", "", -1))
// 		} else if i == 4 {
// 			profile.Weight = util.StringToIntDefault(strings.Replace(string(m[1]), "kg", "", -1))
// 		} else if i == 5 {
// 			profile.Workplace = strings.Replace(string(m[1]), "工作地:", "", -1)
// 		} else if i == 6 {
// 			profile.Income = strings.Replace(string(m[1]), "月收入:", "", -1)
// 		} else if i == 7 {
// 			profile.Occupation = string(m[1])
// 		} else if i == 8 {
// 			profile.Education = string(m[1])
// 		}
// 	}

// 	result := engine.ParserResult{
// 		Items: []interface{}{profile},
// 	}
// 	return result
// }

const proRe = `<div class="des f-cl"[^>]*>([^<]+)</div>`

func ParserProfile(contents []byte, name string, avatar string) engine.ParserResult {
	re := regexp.MustCompile(proRe)
	matchs := re.FindAllSubmatch(contents, -1)

	profile := model.Profile{}
	profile.Name = name
	profile.Avatar = avatar

	for _, m := range matchs {
		body := m[1]
		array := strings.Split(string(body), "|")
		profile.Address = array[0]
		profile.Age = util.StringToIntDefault(strings.Replace(strings.Trim(array[1], " "), "岁", "", -1))
		profile.Education = array[2]
		profile.Marriage = array[3]
		profile.Height = util.StringToIntDefault(strings.Replace(strings.Trim(array[4], " "), "cm", "", -1))
		profile.Income = array[5]
	}

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}
