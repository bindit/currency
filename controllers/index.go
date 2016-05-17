package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type IndexController struct {
	Controller
}

func (ic *IndexController) Init(router *mux.Router) {
	//register routes
	router.HandleFunc("/", ic.index)
}

func (ic *IndexController) index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WELCOME")
}
