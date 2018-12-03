package tcpserver

import (
	"fmt"
	"net"
	"strings"
	"tcpservertest/utils/errutil"
)

type TCPServer struct {
	listener net.Listener
	callback func(string) string
}

func NewServer(hostAddr string, onCmdFunc func(string) string) (tcpserver *TCPServer) {
	var err error
	tcpserver = &TCPServer{
		callback: onCmdFunc,
	}
	if tcpserver.listener, err = net.Listen("tcp", hostAddr); errutil.CheckError(err, "listen error") {
		return
	}
	fmt.Println("Listen, Wait clients..")
	go tcpserver.Serve()
	return tcpserver
}

func (s *TCPServer) Serve() {
	for {
		conn, err := s.listener.Accept()
		if errutil.CheckError(err, "connect accept err ") {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " client connect success")
		go s.connectionHandler(conn)
	}
}

func (s *TCPServer) connectionHandler(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		len, err := conn.Read(buffer)
		if errutil.CheckError(err, "conn read error") {
			return
		}
		var cmds = strings.Split(string(buffer[:len]), "\n")
		fmt.Println(conn.RemoteAddr().String(), " client send:", string(buffer[:len]))
		fmt.Println("get cmds:", cmds)
		for _, cmd := range cmds {
			if cmd != "" {
				if cmd == "quit" {
					conn.Write([]byte("Bye!\n"))
					conn.Close()
					return
				} else {
					var rtnMsg = s.callback(cmd)
					fmt.Println("rtn:", rtnMsg)
					conn.Write([]byte(rtnMsg + "\n"))
				}
			}
		}
	}
}
