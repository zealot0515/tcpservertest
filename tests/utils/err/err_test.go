package err_test

import (
	"fmt"
	"tcpservertest/utils/errutil"
	"testing"
)

func TestErrUtil(t *testing.T) {
	if errutil.CheckError(nil, "") {
		t.Fatal("errutil error")
	}
	if !errutil.CheckError(fmt.Errorf("test err msg"), "test head") {
		t.Fatal("errutil error")
	}
}
