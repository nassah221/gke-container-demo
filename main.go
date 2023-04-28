package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tpl := template.Must(template.ParseFiles("index.html"))

	srv := http.DefaultServeMux
	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request")

		tpl.ExecuteTemplate(w, "index.html", struct {
			Title   string
			BGColor string
		}{
			Title:   "GKE-Demo",
			BGColor: "red",
		})
	})

	log.Println("Listening on :8888")
	http.ListenAndServe(":8888", srv)
}
