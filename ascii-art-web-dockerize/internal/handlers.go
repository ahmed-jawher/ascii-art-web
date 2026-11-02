package internal

import (
	"html/template"
	"net/http"
)

var Templates *template.Template

type PageData struct {
	Result string
	Error  string
	Text   string
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := Templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		HandleInternalError(w, r)
	}
}

func HandleInternalError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	err := Templates.ExecuteTemplate(w, "500.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error  :( ", http.StatusInternalServerError)
	}
}

func HandleBadRequest(w http.ResponseWriter, r *http.Request, message string) {
	w.WriteHeader(http.StatusForbidden)
	err := Templates.ExecuteTemplate(w, "400.html", &PageData{Error: message})
	if err != nil {
		HandleInternalError(w, r)
	}
}
