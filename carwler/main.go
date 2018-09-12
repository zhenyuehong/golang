package main

import (
	"golang/carwler/engine"
	"golang/carwler/scheduler"
	"golang/carwler/zhenai/parse"
)

//提取珍爱网 城市和链接
func main() {
	//engine.ConcurrentEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parse.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})

}
