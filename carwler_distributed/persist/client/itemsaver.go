package client

import (
	"golang/carwler/engine"
	"golang/carwler_distributed/config"
	"golang/carwler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
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
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("item saver err: error saving item %v: %v",
					item, err)
			}
		}

	}()
	return out, nil
}
