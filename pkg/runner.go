package runner

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

const (
	PYTHON = iota
	GOLANG
)

func Run(langauge int, code string) (result string, err error) {
	var cmd *exec.Cmd
	switch langauge {
	case PYTHON:
		cmd = exec.Command("docker", "exec", "python", "python", "main.py")
		err = os.WriteFile("code/main.py", []byte(code), 0644)
	case GOLANG:
		cmd = exec.Command("docker", "exec", "golang", "go", "run", "main.go")
		err = os.WriteFile("code/main.go", []byte(code), 0644)
	default:
		err = errors.New("runner: language not supported")
		return
	}

	if err != nil {
		log.Println("Runner", err)
		return
	}

	output, err := cmd.CombinedOutput()
	result = string(output)
	return
}
