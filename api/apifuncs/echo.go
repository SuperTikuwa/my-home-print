package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if err := exec.Command("lp", "/home/tikuwa/download/hello.txt").Run(); err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.Marshal(PrintStatus{IsFinished: true})

	if err != nil {
		log.Fatal(err, "hoge")
	}

	fmt.Fprintln(w, jsonBytes)

}
