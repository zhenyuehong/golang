package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex //给它加一个锁
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	//但是在加之前要用锁来保护它
	a.value++
	//加完之后要解锁
	defer a.lock.Unlock()
}
func (a *atomicInt) get() int {
	//这里也有锁保护
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
