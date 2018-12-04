package serverinfo

var serverInfoFuncs = map[string]func() interface{}{}

func RegistInfo(infoKey string, infoFunc func() interface{}) {
	if _, ok := serverInfoFuncs[infoKey]; !ok {
		serverInfoFuncs[infoKey] = infoFunc
	}
}

func QueryServerInfo() map[string]interface{} {
	var rtnInfo = map[string]interface{}{}
	for k, v := range serverInfoFuncs {
		rtnInfo[k] = v()
	}
	return rtnInfo
}
