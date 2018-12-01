package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	TCPPort int `json:"tcpport"`
}

var Conf *Config

func init() {
	var bs []byte
	var err error
	if bs, err = ioutil.ReadFile("conf.json"); err != nil {
		fmt.Println("read file error:", err)
		return
	}
	Conf = &Config{}
	if err = json.Unmarshal(bs, Conf); err != nil {
		fmt.Println("config parse error:", err)
		return
	}
}
