package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// db := config.Connect()

	// rows, err := db.Query("SELECT * FROM users")

	// if err != nil {
	// 	// response.Status = 400
	// 	// response.Message = err.Error()
	// 	// SendResponse(w, response.Status, response)

	// 	fmt.Print(err)
	// 	return
	// }

	// var users []models.User
	// var user models.User

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

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		return
	}

	fmt.Println(r.Form.Get("name"))
	fmt.Println(r.Form.Get("password"))

	return
}
