package main

import (
	"fmt"
	"golang/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	//conn, err := net.Dial("tcp", ":1234")
	conn, err := net.Dial("tcp", "golang.org:http")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	fmt.Println(result, err)
}
