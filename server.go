package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	runner "github.com/metallust/compiler-container/pkg"
)

type compileIN struct {
	Code     string `json:"code"`
	Langauge string `json:"language"`
}

func compile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	var in compileIN
	err = json.Unmarshal(body, &in)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	switch in.Langauge {
	case "python":
		code := []byte(in.Code)
		err := os.WriteFile("code/main.py", code, 0644)
		if err != nil {
			log.Fatal("Can't write in code/main.py")
		}
		output, err := runner.Run(runner.PYTHON)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, output)
	default:
		fmt.Fprintln(w, "Langauge not supported...")
	}
}

func main() {

	http.HandleFunc("/", compile)
	err := http.ListenAndServe(":8090", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
