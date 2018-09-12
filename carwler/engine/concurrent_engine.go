package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
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
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("got item %d: %v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadNotifier) {
	//func createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			// tell scheduler i'm ready
			//s.WorkerReady(in)
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
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
