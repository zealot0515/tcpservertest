package serverinfo

var ServerInfos = map[string]int{}

func UpdateInfo(infoKey string, diffValue int, reset bool) {
	if reset {
		ServerInfos[infoKey] = diffValue
	} else {
		ServerInfos[infoKey] += diffValue
	}
}
