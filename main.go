package main

import (
	"fmt"
	"sync"
	"tcpservertest/server/tcpserver"
	"tcpservertest/utils/conf"
)

func main() {
	var hostAddr = fmt.Sprintf("localhost:%d", conf.Conf.TCPPort)
	go tcpserver.Serve(hostAddr)
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("server started")
	wg.Wait()
}
