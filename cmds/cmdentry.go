package cmds

import (
	"fmt"
	"strings"
	"tcpservertest/utils/serverinfo"
	"time"
)

var handlerMap = map[string]func([]string) string{}
var commandCountResetSetting = int64(100000)
var commandCount = int64(0)
var commandCountStartTimespin int64

func CmdEntry(cmd string) string {
	fmt.Println("CmdEntry:", cmd)
	var params = strings.Split(cmd, ",")
	var header = params[0]
	params = params[1:]
	var rtnStr = "cmd not found"
	if handlerFunc, ok := handlerMap[header]; ok {
		rtnStr = handlerFunc(params)
		addCommandCount()
	}
	return rtnStr
}

func RegistCmdHandler(cmdHeader string, handler func([]string) string) {
	handlerMap[cmdHeader] = handler
}

func init() {
	serverinfo.RegistInfo("command per sec", commandPerSecInfo)
	commandCount = commandCountResetSetting
}

func addCommandCount() {
	commandCount++
	if commandCount >= commandCountResetSetting {
		commandCount = 0
		commandCountStartTimespin = time.Now().UTC().Unix()
	}
}

func commandPerSecInfo() interface{} {
	return fmt.Sprintf("%d cmd/s", int(commandCount/(time.Now().UTC().Unix()-commandCountStartTimespin)))
}
