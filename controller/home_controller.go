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

	data := map[string]interface{}{

		"title": "Home | Anonymous Message",
	}

	err = templ.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
