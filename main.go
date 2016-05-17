package main

import (
	"net/http"
	"currency/web"
)

func main() {
	app := web.App{}
	app.Init()

	http.ListenAndServe(":8080", app.Router())
}