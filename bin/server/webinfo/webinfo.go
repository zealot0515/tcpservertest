package webinfo

import (
	"html/template"
	"net/http"
	"sort"
	"tcpservertest/utils/serverinfo"
)

type Info struct {
	Key   string
	Value interface{}
}

type PageData struct {
	Infos []Info
}

func ServeWeb() {
	template := template.Must(template.ParseFiles("webinfo/weblayout.html"))

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		var data = updateInfo()
		template.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

func updateInfo() PageData {
	var infosMap = map[string]Info{}
	var infos = []Info{}
	var keys = []string{}
	for k, v := range serverinfo.QueryServerInfo() {
		keys = append(keys, k)
		infosMap[k] = Info{
			Key:   k,
			Value: v,
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		infos = append(infos, infosMap[k])
	}
	return PageData{
		Infos: infos,
	}
}
