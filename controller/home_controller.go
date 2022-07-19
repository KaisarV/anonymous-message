package controller

import (
	"html/template"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(path.Join("views/homepage", "index.html"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _, name, username, email, _ := validateTokenFromCookies(r)

	data := map[string]interface{}{

		"title":    "Home | Anonymous Message",
		"name":     name,
		"username": username,
		"email":    email,
	}

	err = templ.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
