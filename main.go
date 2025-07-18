package main

import (
	"github.com/MrShanks/templates"
	"log"
	"net/http"
)

func main() {

	component := templates.Layout()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
