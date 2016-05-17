package controllers

import (
	"github.com/gorilla/mux"
)

type IController interface  {
	Init(router *mux.Router)
}

type Controller struct {}