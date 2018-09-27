package main

import (
	"golang/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

//{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
//{"method":"DemoService.Div","params":[{"A":3,"B":0}],"id":1234}
