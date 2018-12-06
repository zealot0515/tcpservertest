package cmds

import (
	"fmt"
	"math"
	"strings"
	"tcpservertest/utils/serverinfo"
	"time"
)

var handlerMap = map[string]func([]string) string{}
var commandCountResetSetting = int64(100000)
var commandCount = int64(0)
var pendingCount = int64(0)
var finishCommandCount = int64(0)
var commandCountStartTimespin int64

func CmdEntry(cmd string) string {
	fmt.Println("CmdEntry:", cmd)
	var params = strings.Split(cmd, ",")
	var header = params[0]
	params = params[1:]
	var rtnStr = "cmd not found"
	if handlerFunc, ok := handlerMap[header]; ok {
		beforeCommand()
		rtnStr = handlerFunc(params)
		afterCommand()
	}
	return rtnStr
}

func RegistCmdHandler(cmdHeader string, handler func([]string) string) {
	handlerMap[cmdHeader] = handler
}

func init() {
	commandCount = commandCountResetSetting
	pendingCount = 0
	finishCommandCount = 0
	serverinfo.RegistInfo("command per sec", func() interface{} {
		return fmt.Sprintf("%d cmd/s", int(commandCount/(time.Now().UTC().Unix()-commandCountStartTimespin)))
	})
	serverinfo.RegistInfo("finish commands", func() interface{} {
		return fmt.Sprintf("%d cmds", finishCommandCount)
	})
	serverinfo.RegistInfo("pending cmds", func() interface{} {
		return fmt.Sprintf("%d cmds", pendingCount)
	})
}

func beforeCommand() {
	pendingCount++
}

func afterCommand() {
	pendingCount--
	finishCommandCount++
	if finishCommandCount >= math.MaxInt64 {
		finishCommandCount = 0
	}
	addCommandCount()
}

func addCommandCount() {
	commandCount++
	if commandCount >= commandCountResetSetting {
		commandCount = 0
		commandCountStartTimespin = time.Now().UTC().Unix()
	}
}
