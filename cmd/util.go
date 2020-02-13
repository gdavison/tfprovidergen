package cmd

import (
	"fmt"
	"os"
)

func errorExitf(format string, a ...interface{}) {
	errorExit(fmt.Sprintf(format, a...))
}

func errorExit(a interface{}) {
	fmt.Println(a)
	os.Exit(1)
}
