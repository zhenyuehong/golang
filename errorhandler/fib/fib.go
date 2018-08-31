package fib

//定义一个斐波那契数列的生成器
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
