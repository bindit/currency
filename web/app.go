package web

import (
	"github.com/gorilla/mux"
	"currency/controllers"
)

type App struct {
	router *mux.Router
}

func (app *App) Init() {
	app.initRouter()
	app.initControllers()
}

func (app *App) initRouter() {
	app.router = mux.NewRouter()
}

func (app *App) initControllers() {
	indexController := controllers.IndexController{}
	indexController.Init(app.router)

	currencyController := controllers.CurrencyController{}
	currencyController.Init(app.router)
}

func (app *App) Router() *mux.Router {
	return app.router
}

