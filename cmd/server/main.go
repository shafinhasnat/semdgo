package main

import (
	"log"
	"net/http"

	"github.com/shafinhasnat/semdgo/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.Handle)
	log.Println("Starting server on port 80")
	http.ListenAndServe(":80", nil)
}
