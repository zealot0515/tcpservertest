package cmds

import (
	"fmt"
	"strings"
	"tcpservertest/utils/serverinfo"
)

var handlerMap = map[string]func([]string) string{}

func CmdEntry(cmd string) string {
	fmt.Println("CmdEntry:", cmd)
	var params = strings.Split(cmd, ",")
	var header = params[0]
	params = params[1:]
	var rtnStr = "cmd not found"
	if handlerFunc, ok := handlerMap[header]; ok {
		rtnStr = handlerFunc(params)
		serverinfo.UpdateInfo("cmdExecCount", 1, false)
	}
	return rtnStr
}

func RegistCmdHandler(cmdHeader string, handler func([]string) string) {
	handlerMap[cmdHeader] = handler
}
