package persist

import (
	"context"
	"errors"
	"golang/carwler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		//must turn off sniff in docker
		elastic.SetSniff(false), //维护集群的状态
		//elastic.SetURL() //这个可以省略，省略即默认9200端口
	)
	if err != nil {
		return nil, err
	}

	//item saver 将通过 OUT 实现和engine直接进行传递
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d: %v", itemCount, item)
			itemCount++
			err := SaveItem(client, index, item)
			if err != nil {
				log.Printf("item saver err: error saving item %v: %v",
					item, err)
			}
		}

	}()
	return out, nil
}

func SaveItem(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	//client.Index()  用来创建或者修改数据
	indexService := client.Index().
		Index(index).
		Type(item.Type).Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
