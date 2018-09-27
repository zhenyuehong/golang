package main

import (
	"golang/carwler/engine"
	"golang/carwler/model"
	"golang/carwler_distributed/config"
	"golang/carwler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	host := ":1234"
	//start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//call SaverItem
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1434875416",
		Type: "zhenai",
		Id:   "1434875416",
		Payload: model.Profile{
			Name:       "提线木偶",
			Gender:     "男",
			Age:        29,
			Height:     168,
			Weight:     "--",
			Income:     "3001-5000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "运营管理",
			Hukou:      "辽宁鞍山",
			Xingzuo:    "天蝎座",
			House:      "--",
			Car:        "未购车",
		},
	}
	result := ""
	//client.Call("ItemSaverService.SaveItem", item, &result)
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
