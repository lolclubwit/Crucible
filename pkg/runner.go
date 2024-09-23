package runner

import (
	"errors"
	"fmt"
	"os/exec"
)

const (
	PYTHON = iota
	GOLANG
)

func Run(langauge int) (result string, err error) {
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	var cmd *exec.Cmd
	switch langauge {
	case PYTHON:
		cmd = exec.Command("docker", "start", "-i", "python")
	case GOLANG:
		cmd = exec.Command("docker", "start", "-i", "golang")
	default:
		err = errors.New("language not supported")
		return
	}

	output, err := cmd.Output()
	if err != nil {
		return
	}

	result = string(output)
	return
}
