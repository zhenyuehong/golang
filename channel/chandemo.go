package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}

func work(id int, c chan int) {
	//for {
	//	n, ok := <-c  //ok判断channel里面的数据是不是空
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %d\n",
	//		id, n)
	//}

	//另一种做法
	for n := range c {
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func chanDemo() {
	//c := make(chan int)
	////给channel发数据的时候，要有goroutine来接收channel的数据，否则就会deadlock
	//go worker(c)
	//c <- 1
	//c <- 2

	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //3是给channel 3个缓冲区，这样sent 3个数据就不会deal lock，超过3个才会deal lock
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)

}

func channelClose() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
