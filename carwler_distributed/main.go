package main

import (
	"flag"
	"golang/carwler/engine"
	"golang/carwler/scheduler"
	"golang/carwler/zhenai/parse"
	itemSaverClient "golang/carwler_distributed/persist/client"
	"golang/carwler_distributed/rpcsupport"
	workerClient "golang/carwler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver_host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)") //逗号分离
)

//分布式  抽离ItemSaver
//提取珍爱网 城市和链接
func main() {
	//开client之前先把server启动
	//carwler_distributed/persist/server/main.go
	//itemSaver, err := itemSaverClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	flag.Parse()
	itemSaver, err := itemSaverClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := workerClient.CreateProcessor(pool)

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

func createClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s : %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		//轮流顺序分发
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
