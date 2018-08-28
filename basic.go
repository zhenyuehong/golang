package main

import "fmt"

func main() {
	constant()
}

func constant() {

	const (
		name = "abc"
		a    = 4
		b    = 5
		c    = 7
	)
	fmt.Printf("%s,%d,%d,%d", name, a, b, c)
}
