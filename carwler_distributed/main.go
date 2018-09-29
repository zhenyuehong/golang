package main

import (
	"fmt"
	"golang/carwler/engine"
	"golang/carwler/scheduler"
	"golang/carwler/zhenai/parse"
	"golang/carwler_distributed/config"
	itemSaverClient "golang/carwler_distributed/persist/client"
	workerClient "golang/carwler_distributed/worker/client"
)

//分布式  抽离ItemSaver
//提取珍爱网 城市和链接
func main() {
	//开client之前先把server启动
	//carwler_distributed/persist/server/main.go
	itemSaver, err := itemSaverClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := workerClient.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parse.ParseCityList, "ParseCityList"),
		//ParserFunc: parse.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parse.ParseCity,
	//})
}
