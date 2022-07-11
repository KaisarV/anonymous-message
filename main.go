package main

import (
	"log"
	"net/http"
	"ngl-link/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomeHandler)
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
