package webinfo

import (
	"html/template"
	"net/http"
	"tcpservertest/utils/serverinfo"
)

type Info struct {
	Key   string
	Value int
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
	var infos = []Info{}
	for k, v := range serverinfo.ServerInfos {
		infos = append(infos, Info{
			Key:   k,
			Value: v,
		})
	}
	return PageData{
		Infos: infos,
	}
}
