package main

import (
	"net/http"

	"www.tikuwa.monster/home-server/apifuncs"
)

func main() {
	http.HandleFunc("/p-api", apifuncs.PrintHandler)

	http.ListenAndServe(":1234", nil)
}
