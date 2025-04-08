package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./content"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
