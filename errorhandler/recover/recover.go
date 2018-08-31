package main

import (
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred: ", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is a error"))
	b := 0
	a := 5 / b
	fmt.Println(a)
}
