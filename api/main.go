package main

import (
	"net/http"

	"www.tikuwa.monster/home-server/apifuncs"
)

func main() {
	http.HandleFunc("/echo", apifuncs.EchoHandler)

	http.ListenAndServe(":1234", nil)
}
