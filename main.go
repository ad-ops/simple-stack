package main

import (
	// "database/sql"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	// "os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"github.com/ad-ops/simple-stack/handlers"
)

func generatePageRoutes() {
	layouts, err := filepath.Glob("web/layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}

	pages, err := filepath.Glob("web/pages/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {
		fileName := filepath.Base(page)
		pathName := "/" + strings.Replace(fileName, ".html", "", 1)
		files := append(layouts, page)
		template := template.Must(template.ParseFiles(files...))

		http.HandleFunc(pathName, func(w http.ResponseWriter, r *http.Request) {
			log.Println("got request for " + pathName)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			err := template.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}
}

func main() {
	log.Println("Hello, world!")

	generatePageRoutes()

	fs := http.FileServer(http.Dir("web/public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	
	http.HandleFunc("/stuff", handlers.StuffHandler)
	http.HandleFunc("/available-pets", handlers.AvailablePetsHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	})

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
