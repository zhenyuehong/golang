package main

//函数
import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(apply(pow, 3, 4))
	//另一种写法
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

}

func apply(op func(int, int) int, a, b int) int {
	//使用反射的写法获取到func name
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling func %s with args (%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//循环，求和
func sum(number ...int) int {
	s := 0
	for i := range number {
		s += number[i]
	}
	return s
}
