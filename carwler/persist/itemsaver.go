package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	//item saver 将通过 OUT 实现和engine直接进行传递
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d: %v", itemCount, item)
			itemCount++
			saveItem(item)
		}

	}()
	return out
}

func saveItem(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		//must turn off sniff in docker
		elastic.SetSniff(false), //维护集群的状态
		//elastic.SetURL() //这个可以省略，省略即默认9200端口
	)
	if err != nil {
		return "", err
	}

	//client.Index()  用来创建或者修改数据
	resp, err := client.Index().
		Index("dating_profile").Type("zhenai").
		BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Printf("%+v/n", resp)
	return resp.Id, nil
}
