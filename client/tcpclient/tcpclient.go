package tcpclient

import (
	"fmt"
	"net"
	"tcpservertest/utils/errutil"
)

type TCPClient struct {
	conn net.Conn
}

func Connect(serverAddr string) (client *TCPClient) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	if errutil.CheckError(err, "reslove addr error") {
		return nil
	}

	client = &TCPClient{}
	client.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if errutil.CheckError(err, "dail tcp error") {
		return nil
	}

	fmt.Println("connect ", serverAddr, " success!")
	return client
}

func (c *TCPClient) Send(msg string) {
	c.conn.Write([]byte(msg))
	fmt.Println("send ", msg, ", done!")
}
