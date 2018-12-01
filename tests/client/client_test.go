package clinet_test

import (
	"fmt"
	"sync"
	"tcpservertest/client/tcpclient"
	"tcpservertest/utils/conf"
	"testing"
)

func TestClient(t *testing.T) {
	var serverAddr = fmt.Sprintf("localhost:%d", conf.Conf.TCPPort)
	var wg sync.WaitGroup
	var client = tcpclient.Connect(serverAddr, func(receiveMsg string) {
		fmt.Println("Get ServerMsg:", receiveMsg)
		wg.Done()
	})
	if client != nil {
		for i := 0; i < 10; i++ {
			client.Send("Hello server!\n")
			wg.Add(1)
		}
	} else {
		t.Fatal("can't get clinet, check connectioin")
	}
	wg.Wait()
	fmt.Println("Get All MsgBack, done!")
}
