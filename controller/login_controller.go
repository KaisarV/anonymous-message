package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ngl-link/config"
	model "ngl-link/models"
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

	db := config.Connect()

	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	rows, err := db.Query("SELECT * FROM  users WHERE email = ? AND password = ?", email, password)

	if err != nil {
		fmt.Println("Error")
		return
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Age, &user.Usertype); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return
}
