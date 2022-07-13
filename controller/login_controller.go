package controller

import (
	"html/template"
	"net/http"
	"path"
)

func Login(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(path.Join("views/auth", "login.html"), path.Join("views/layout", "layout.html"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{

		"title": "Login | Anonymous Message",
	}

	err = templ.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
