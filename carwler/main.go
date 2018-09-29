package main

import (
	"golang/carwler/engine"
	"golang/carwler/persist"
	"golang/carwler/scheduler"
	"golang/carwler/zhenai/parse"
	"golang/carwler_distributed/config"
)

//提取珍爱网 城市和链接
func main() {
	//engine.ConcurrentEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parse.ParseCityList,
	//})

	itemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		//ParserFunc: parse.ParseCityList,
		Parser: engine.NewFuncParser(parse.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parse.ParseCity,
	//})
}
