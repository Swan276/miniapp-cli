package utils

import (
	"fmt"
	"os"
)

func Success(msg string) {
	fmt.Printf("\033[32m%s \033[0;30m%s\n", "✔️", msg)
}

func ErrorMsg(msg string) {
	fmt.Printf("\033[1;31m%s\033[0m\n", msg)
}

func Abort(msg string) {
	ErrorMsg(msg)
	os.Exit(1)
}

func Info(msg string) {
	fmt.Printf("\033[1;34m%s\033[0m\n", msg)
}

func Warning(msg string) {
	fmt.Printf("\033[1;33m%s\033[0m\n", msg)
}
