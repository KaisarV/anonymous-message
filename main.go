package main

import (
	"log"
	"net/http"
	"ngl-link/controller"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", controller.Authenticate(controller.HomeHandler, 1))
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/process", controller.ProcessLogin)
	router.HandleFunc("/logout", controller.Logout)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
