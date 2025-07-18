package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrShanks/Tick/static/templates"
	"github.com/a-h/templ"
)

func main() {

	index := templates.Layout("Tick")

	http.Handle("/", templ.Handler(index))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
