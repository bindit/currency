package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"currency/controllers"
)

func main() {
	router := mux.NewRouter()

	indexController := controllers.IndexController{}
	indexController.Init(router)

	http.ListenAndServe(":8080", router)
}