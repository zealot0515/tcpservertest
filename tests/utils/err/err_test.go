package err_test

import (
	"fmt"
	"tcpservertest/utils/errutil"
	"testing"
)

func TestErrUtil(t *testing.T) {
	if errutil.CheckError(nil, "") {
		fmt.Println("errutil error")
	}
	if !errutil.CheckError(fmt.Errorf("test err msg"), "test head") {
		fmt.Println("errutil error")
	}
}
