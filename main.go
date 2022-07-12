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

	router.HandleFunc("/", controller.HomeHandler)
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
