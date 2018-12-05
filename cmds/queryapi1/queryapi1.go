package queryapi1

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"tcpservertest/client/httprequest"
	"tcpservertest/cmds"
	"tcpservertest/utils/conf"
	"tcpservertest/utils/errutil"
	"tcpservertest/utils/serverinfo"
	"tcpservertest/utils/timeutil"
	"time"
)

var r *rand.Rand
var apiHost string
var lastExecTime = int64(-1)

func queryApi1(params []string) string {
	var tracker = timeutil.TimeTracker{}
	tracker.Begin()
	var req = &httprequest.ReqParams{
		Method: "GET",
		URI:    apiHost,
		Body:   "",
		QueryParams: map[string]string{
			"randomstring": fmt.Sprintf("%d", r.Int63n(math.MaxInt64)),
		},
	}
	body, err := httprequest.SendRequest(req)
	if errutil.CheckError(err, "send http req err") {
		return "fail"
	}
	var result = map[string]string{}
	err = json.Unmarshal(body, &result)
	if errutil.CheckError(err, "api unmarshal err") {
		return "fail"
	}
	tracker.End()
	logApiResponseTime(tracker)
	return result["returnstring"]
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UTC().Unix()))
	cmds.RegistCmdHandler("queryapi1", queryApi1)
	registCmdApiInfoFunc()
	apiHost = "http://" + conf.Conf.APIHost + "/v1/demo/demoapi"
}

func logApiResponseTime(tracker timeutil.TimeTracker) {
	if lastExecTime > 0 {
		lastExecTime = (lastExecTime + tracker.GetDurationTime()) / 2
	} else {
		lastExecTime = tracker.GetDurationTime()
	}
}

func registCmdApiInfoFunc() {
	serverinfo.RegistInfo("Cmd_queryApi1_execTime", func() interface{} {
		if lastExecTime > 0 {
			return fmt.Sprintf("%d us", lastExecTime)
		} else {
			return "--"
		}
	})
}
