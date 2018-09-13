package engine

//单任务版 engine
import (
	"fmt"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//print url
		//log.Printf("fetching url: %s", r.Url)
		//body, err := fetcher.Fetch(r.Url)
		//if err != nil {
		//	log.Printf("fetcher error: fether url %s, %v", r.Url, err)
		//	continue //继续处理下一个request
		//}
		//parserResult := r.ParserFunc(body)

		parserResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			fmt.Printf("got item %v\n", item)
		}
	}
}
