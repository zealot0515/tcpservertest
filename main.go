package main

import (
	"fmt"
	"sync"
	"tcpservertest/cmds"
	"tcpservertest/server/tcpserver"
	"tcpservertest/utils/conf"
)

func main() {
	var hostAddr = fmt.Sprintf("localhost:%d", conf.Conf.TCPPort)
	var server = tcpserver.NewServer(hostAddr, cmds.CmdEntry)
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("server started:", server)
	wg.Wait()
}
