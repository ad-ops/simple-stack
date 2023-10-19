package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type PetData struct {
	Name string
	Age  int
}

func AvailablePetsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for available pets")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.ParseFiles("web/htmx/available-pets.html"))

	pets := []PetData{
		{Name: "Fido", Age: 3},
	}
	err := tmpl.Execute(w, pets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
