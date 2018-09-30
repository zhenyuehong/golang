package main

import (
	"flag"
	"fmt"
	"golang/carwler_distributed/rpcsupport"
	"golang/carwler_distributed/worker"
	"log"
)

//命令行参数
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
