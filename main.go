package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	//第一个版本
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:       "http://www.zhenai.com/zhenghun",
	// 	ParserFun: parser.ParserCityList,
	// })

	//第二个版本
	//限制10个城市 /zhenai/parser/citylist
	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 100,
	// }
	// e.Run(engine.Request{
	// 	Url:       "http://www.zhenai.com/zhenghun",
	// 	ParserFun: parser.ParserCityList,
	// })

	//第三个版本
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParserCityList,
	})

}
