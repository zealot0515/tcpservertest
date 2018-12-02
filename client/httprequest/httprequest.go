package httprequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"tcpservertest/utils/errutil"
)

type ReqParams struct {
	Method      string
	URI         string
	Body        string
	QueryParams map[string]string
}

func SendRequest(req *ReqParams) (body []byte, err error) {
	if req.URI != "" {
		httpClient := &http.Client{}
		urlParam := url.Values{}
		for k, v := range req.QueryParams {
			urlParam.Add(k, v)
		}
		requestURI := fmt.Sprintf("%s?%s", req.URI, urlParam.Encode())
		var request *http.Request
		request, err = http.NewRequest(req.Method, requestURI, strings.NewReader(req.Body))
		if errutil.CheckError(err, "http request create err") {
			return
		}
		if req.Method == "POST" || req.Method == "PUT" {
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		}

		var response *http.Response
		response, err = httpClient.Do(request)

		if errutil.CheckError(err, "http request do err") {
			return
		}
		defer response.Body.Close()
		body, err = ioutil.ReadAll(response.Body)
		if errutil.CheckError(err, "read response.body err") {
			return
		}
	}
	return
}
