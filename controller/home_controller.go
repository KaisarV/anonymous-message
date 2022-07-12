package controller

import (
	"html/template"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(path.Join("views", "index.html"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
