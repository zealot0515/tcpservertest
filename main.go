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
	"time"
)

func main() {
	var hostAddr = fmt.Sprintf("localhost:%d", conf.Conf.TCPPort)
	var server = tcpserver.NewServer(hostAddr, cmds.CmdEntry)
	var wg sync.WaitGroup
	go webinfo.ServeWeb()
	wg.Add(1)
	fmt.Println("server started:", server)
	go func() {
		ticker := time.NewTicker(1000 * time.Millisecond)
		defer ticker.Stop()
		for {
			<-ticker.C
			serverinfo.UpdateInfo("sessionCount", server.SessionCount(), true)
			serverinfo.UpdateInfo("cmdExecCount", 0, true)
		}
	}()
	wg.Wait()
}
