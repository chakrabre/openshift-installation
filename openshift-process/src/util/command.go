package util

import (
	"os/exec"
	"os"
)

func Execute(command string, arg ...string) error {
	return ExecuteDir("", command, arg...)
}

func ExecuteDir(dir string, command string, arg ...string) error {
	cmd := exec.Command(command, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = dir
	err := cmd.Run()

	return err
}