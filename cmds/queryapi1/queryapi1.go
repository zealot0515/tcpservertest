package queryapi1

import "tcpservertest/cmds"

func queryApi1(params []string) string {
	return "bbb"
}

func init() {
	cmds.RegistCmdHandler("queryapi1", queryApi1)
}
