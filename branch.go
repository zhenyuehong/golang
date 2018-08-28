package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const fileName = "abc.txt"

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("content: %s\n", content)

}
