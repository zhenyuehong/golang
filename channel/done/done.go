package main

import (
	"fmt"
	"sync"
)

func main() {
	chanDemo()

}

func doWork(id int, w worker) {
	//通过通信来共享内存
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		//go func() {done <- true}()//要一口气print完两个任务就要另开一个goroutine来做
		//done <- true
		w.done()
	}
}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: make(chan bool),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//1、这样，就是按照顺序来打印了，这就是等一个任务结束，下一个任务开始
	//for i := 0; i < 10; i++ {
	//	workers[i].in <- 'a' + i
	//	<- workers[i].done
	//}
	//for i := 0; i < 10; i++ {
	//	workers[i].in <- 'A' + i
	//	<- workers[i].done
	//}

	//2、我们需要一口气结束 wait for all of them
	//for i, worker := range workers {
	//	worker.in <- 'a' + i
	//}
	//for i, worker := range workers {
	//	worker.in <- 'A' + i
	//}
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}

	//3、逐个结束打印任务
	//for i, worker := range workers {
	//	worker.in <- 'a' + i
	//}
	//for _, worker := range workers {
	//	<-worker.done
	//}
	//for i, worker := range workers {
	//	worker.in <- 'A' + i
	//}
	//for _, worker := range workers {
	//	<-worker.done
	//}

	//使用waitGroup来达到2的效果
	for i, worker := range workers {
		worker.in <- 'a' + i
		wg.Add(1)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		wg.Add(1)
	}
	wg.Wait()

}
