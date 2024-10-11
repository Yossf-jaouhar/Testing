package test

import (
	"log"
	"net/http"
	"text/template"
)

func HomHndler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("Templates/hh.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	
	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
	}
}
