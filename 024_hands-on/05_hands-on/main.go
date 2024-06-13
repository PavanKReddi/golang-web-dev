package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dogs)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Template execution failed: %v", err)
	}
}
