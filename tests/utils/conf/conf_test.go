package conf_test

import (
	"tcpservertest/utils/conf"
	"testing"
)

func TestConf(t *testing.T) {
	if conf.Conf.TCPPort != 6666 {
		t.Fatal("config error")
	}
}
