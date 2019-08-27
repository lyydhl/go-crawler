package main

import (
	"go-crawler/engine"
	"go-crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParserCityList,
	})

	// all, err := fetcher.Fetch("https://album.zhenai.com/u/1445559021")
	// if err != nil {
	// 	panic(err)
	// }

	// parser.ParserProfile1(all, "aa")

}
