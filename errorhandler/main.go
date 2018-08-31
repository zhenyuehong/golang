package main

import (
	"bufio"
	"example/errorhandler/fib"
	"fmt"
	"os"
)

func main() {
	writeFile("errorhandler/fib.txt")
	//writeFile("fib.txt")
}
func writeFile(fileName string) {
	//file, err := os.Create(fileName)
	//if err != nil {
	//	panic(err)
	//}

	//error  handler
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s ,%s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}
