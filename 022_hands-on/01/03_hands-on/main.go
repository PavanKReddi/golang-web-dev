package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("template.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "root")
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "woof woof")
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Greeting string
	}{
		Greeting: "Hello Human",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
