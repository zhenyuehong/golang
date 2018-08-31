package main

import (
	"bufio"
	"example/errorhandler/fib"
	"fmt"
	"io"
	"strings"
)

//利用函数的闭包来打印斐波那契数列
//斐波那契数列  第三个数的前面两个数的和
func main() {
	var f intGen = fib.Fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	printFileContents(f)

}

////定义一个斐波那契数列的生成器
//func fibonacci() intGen {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}

type intGen func() int

//为函数实现接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//todo incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

//从文件里面读取内容并打印，这里是把 intGen 当成一个文件来读
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
