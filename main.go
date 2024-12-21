package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	httpDir := http.Dir(argsWithoutProg[0])

	fs := http.FileServer(httpDir)

	http.Handle("/", fs)

	fmt.Println("serving folder: " + argsWithoutProg[0] + " at port :6464")

	http.FileServer(httpDir)

	http.ListenAndServe(":6464", nil)

}
