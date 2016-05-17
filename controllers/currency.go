package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"currency/model"
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
	c := model.NewCurrency("USD", 1)

	fmt.Printf("PRINT LIST OF CURRENCIES: %s - converter: %f\n", c.GetName(), c.GetConverter())
}

func (cc *CurrencyController) exchange(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CHANGE")
}

