package utils

import (
	"os"
	"os/exec"
	"strings"
)

func ParseCommand(cmd string) []string {
	return strings.Split(cmd, " ")
}

func RunCommand(commandArgs []string) error {
	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)

	cmd.Stdout = os.Stdout

	return cmd.Run()
}
