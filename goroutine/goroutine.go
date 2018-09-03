package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)

	//var a [10]int
	//for i:=0; i<10;i++  {
	//	go func(i int) {
	//		for{
	//			a[i]++
	//			//若不主动交出控制器，就会霸占CPU，程序就永远在这里面死循环，原因是没有机会交出控制权
	//
	//			runtime.Gosched()//协程主动交出控制权，让别人也有机会运行
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
