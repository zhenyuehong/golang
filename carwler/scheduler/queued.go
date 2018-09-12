package scheduler

import "golang/carwler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request //chan engine.Request 其实是worker类型，而worker是chan engine.Request类型
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//它负责告诉我们，外界已经有一个worker 已经ready了，可以负责接收request
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (*QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		//new two queue 来对 r and w 进行任务分配
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r) //send r to a worker
			case w := <-s.workerChan:
				workerQ = append(workerQ, w) //send next request to w
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
