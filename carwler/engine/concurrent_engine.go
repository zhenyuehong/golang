package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(r Request) (ParseResult, error)

type Scheduler interface {
	ReadNotifier
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	WorkerChan() chan Request //workerChan由 scheduler 去决定
	//WorkerReady(chan Request)
	Run()
}
type ReadNotifier interface {
	WorkerReady(chan Request)
}

//queue
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("got item %d: %v", itemCount, item)
			//itemCount++

			//save(item)  这样不行，这个要马上脱手，不能做复杂的io操作
			//go save(item) 这样可行，可以让他在后台执行save 操作
			//go func() {save(itemChan <- item)}()  这样也可行，用goroutine的方法执行操作,我们将采用这种方法进行操作
			//保存 item 数据
			go func() {
				e.ItemChan <- item
			}()

		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadNotifier) {
	//func createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			// tell scheduler i'm ready
			//s.WorkerReady(in)
			ready.WorkerReady(in)
			request := <-in
			//result, err := Worker(request)
			result, err := e.RequestProcessor(request) //call rpc
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//去重
var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

//no queue
//func (e *ConcurrentEngine) Run(seeds ...Request) {
//
//	in := make(chan Request)
//	out := make(chan ParseResult)
//
//	e.Scheduler.ConfigureMasterWorkerChan(in)
//
//	for i := 0; i < e.WorkerCount; i++ {
//		createWorker(in, out)
//	}
//
//	for _, r := range seeds {
//		e.Scheduler.Submit(r)
//	}
//
//	itemCount := 0
//	for {
//		result := <-out
//		for _, item := range result.Items {
//			log.Printf("got item %d: %v", itemCount, item)
//			itemCount++
//		}
//		for _, request := range result.Requests {
//			e.Scheduler.Submit(request)
//		}
//	}
//}
//
//func createWorker(in chan Request, out chan ParseResult) {
//	go func() {
//		for {
//			 // tell scheduler i'm ready
//			request := <-in
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result
//		}
//	}()
//}
