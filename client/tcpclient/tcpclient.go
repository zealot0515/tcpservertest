package tcpclient

import (
	"fmt"
	"net"
	"strings"
	"tcpservertest/utils/errutil"
)

type TCPClient struct {
	conn            net.Conn
	receiveCallback func(string)
}

func Connect(serverAddr string, callback func(string)) (client *TCPClient) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	if errutil.CheckError(err, "reslove addr error") {
		return nil
	}

	client = &TCPClient{
		receiveCallback: callback,
	}
	client.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if errutil.CheckError(err, "dail tcp error") {
		return nil
	}

	fmt.Println("connect ", serverAddr, " success!")
	go client.waitReceive()
	return client
}

func (c *TCPClient) Send(msg string) {
	c.conn.Write([]byte(msg))
}

func (c *TCPClient) waitReceive() {
	buffer := make([]byte, 2048)
	for {
		len, err := c.conn.Read(buffer)
		if errutil.CheckError(err, "client conn read error") {
			return
		}
		var receiveMsgs = strings.Split(string(buffer[:len]), "\n")
		for _, msg := range receiveMsgs {
			if msg != "" {
				c.receiveCallback(msg)
			}
		}
	}
}
