package queryapi1

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"tcpservertest/client/httprequest"
	"tcpservertest/cmds"
	"tcpservertest/utils/errutil"
	"time"
)

var r *rand.Rand

func queryApi1(params []string) string {
	var req = &httprequest.ReqParams{
		Method: "GET",
		URI:    "http://127.0.0.1:7777/v1/demo/demoapi",
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
	return result["returnstring"]
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UTC().Unix()))
	cmds.RegistCmdHandler("queryapi1", queryApi1)
}
