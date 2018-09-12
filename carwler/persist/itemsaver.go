package persist

import "log"

func ItemSaver() chan interface{} {
	//item saver 将通过 OUT 实现和engine直接进行传递
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d: %v", itemCount, item)
			itemCount++
		}

	}()
	return out
}
