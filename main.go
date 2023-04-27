package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.DefaultServeMux
	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})

	log.Println("Listening on :8888")
	http.ListenAndServe(":8888", srv)
}
