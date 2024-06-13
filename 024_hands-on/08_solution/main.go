package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", dogsHandler)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))

	log.Printf("Starting server on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func dogsHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
