package main

import (
	"fmt"
	"golang/carwler/engine"
	"golang/carwler/scheduler"
	"golang/carwler/zhenai/parse"
	"golang/carwler_distributed/config"
	"golang/carwler_distributed/persist/client"
)

//分布式  抽离ItemSaver
//提取珍爱网 城市和链接
func main() {

	itemSaver, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
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
