package main

import (
	"fmt"
	"golang/carwler_distributed/config"
	"golang/carwler_distributed/rpcsupport"
	"golang/carwler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
