package main

import (
	"golang/carwler/engine"
	"golang/carwler/persist"
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
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parse.ParseCity,
	//})
}
