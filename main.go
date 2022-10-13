package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println("Starting server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received on home path")
		fmt.Fprintf(w, "Welcome to sample go pprof server!")
	})
	log.Println(http.ListenAndServe(":6060", nil))
}
