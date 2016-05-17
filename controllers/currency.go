package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type CurrencyController struct {
	Controller
}

func (cc *CurrencyController) Init(router *mux.Router) {
	//register routes
	router.HandleFunc("/list", cc.list)
	router.HandleFunc("/exchange", cc.exchange)
}

func (cc *CurrencyController) list(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PRINT LIST OF CURRENCIES")
}

func (cc *CurrencyController) exchange(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CHANGE")
}

