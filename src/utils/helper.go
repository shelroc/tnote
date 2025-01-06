package utils

import (
	"fmt"
	"os"
)

func HandleError(e error) {
	if e != nil {
		msg := fmt.Errorf("%v", e)
		fmt.Println(msg)
		os.Exit(1)
	}
}
