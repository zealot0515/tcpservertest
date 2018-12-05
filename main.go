package main

import (
	"fmt"
	"sync"
	"tcpservertest/cmds"
	_ "tcpservertest/cmds/queryapi1"
	"tcpservertest/server/tcpserver"
	"tcpservertest/utils/conf"
	"tcpservertest/utils/serverinfo"
	"tcpservertest/webinfo"
)

func main() {
	var hostAddr = fmt.Sprintf("0.0.0.0:%d", conf.Conf.TCPPort)
	var server = tcpserver.NewServer(hostAddr, cmds.CmdEntry)
	serverinfo.RegistInfo("sessionCount", func() interface{} {
		return server.SessionCount()
	})
	var wg sync.WaitGroup
	go webinfo.ServeWeb()
	wg.Add(1)
	fmt.Println("server started:", server)
	wg.Wait()
}
