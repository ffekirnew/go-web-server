package main

import (
	"encoding/json"
	"net/http"
)

type FormSchema struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	http.ServeFile(w, r, "./static/form.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	formData := &FormSchema{}
	formData.Name = r.FormValue("name")
	formData.Email = r.FormValue("email")
	formData.Message = r.FormValue("message")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(formData); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
