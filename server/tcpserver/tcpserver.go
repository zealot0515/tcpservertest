package tcpserver

import (
	"fmt"
	"net"
	"tcpservertest/utils/errutil"
)

var listener net.Listener

func Serve(hostAddr string) {
	var err error
	if listener, err = net.Listen("tcp", hostAddr); errutil.CheckError(err, "listen error") {
		return
	}
	fmt.Println("Listen, Wait clients..")
	for {
		conn, err := listener.Accept()
		if errutil.CheckError(err, "connect accept err ") {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " client connect success")
		connectionHandler(conn)
	}
}

func connectionHandler(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		len, err := conn.Read(buffer)
		if errutil.CheckError(err, "conn read error") {
			return
		}
		fmt.Println(conn.RemoteAddr().String(), " client send:", string(buffer[:len]))
	}
}
