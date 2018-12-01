package errutil

import (
	"fmt"
)

func CheckError(err error, head string) bool {
	if err != nil {
		fmt.Println("[Error]", head, ":", err)
		return true
	}
	return false
}
