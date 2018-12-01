package server_test

import (
	"tcpservertest/server/tcpserver"
	"testing"
)

func TestServer(t *testing.T) {
	var server = tcpserver.NewServer("127.0.0.1:1234", func(getCmd string) string {
		return "callback"
	})
	if server == nil {
		t.Fatal("can't create tcpserver")
	}
}
