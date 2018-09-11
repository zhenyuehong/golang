package engine

//并发版 engine
import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	for {
		result := <-out
		for _, items := range result.Items {
			fmt.Printf("got item: %v", items)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
