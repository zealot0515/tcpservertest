package clinet_test

import (
	"fmt"
	"tcpservertest/client/tcpclient"
	"tcpservertest/utils/conf"
	"testing"
)

func TestClient(t *testing.T) {
	var serverAddr = fmt.Sprintf("localhost:%d", conf.Conf.TCPPort)
	var client = tcpclient.Connect(serverAddr)
	if client != nil {
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
		client.Send("Hello server!")
	}
}
