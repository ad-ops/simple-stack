package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type StuffData struct {
	Stuff string
}

func StuffHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for stuff")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.ParseFiles("web/htmx/stuff.html"))

	stuffData := &StuffData{Stuff: "this is the stuff!"}
	err := tmpl.Execute(w, stuffData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
