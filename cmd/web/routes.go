package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes(cfg *config) *httprouter.Router {

	router := httprouter.New()
	router.GET("/", app.home)
	router.GET("/hello/:name", app.home)
	router.GET("/snippet", app.showSnippet)
	router.POST("/snippet/create", app.createSnippet)

	router.ServeFiles("/static/*filepath", http.Dir(cfg.staticfiles))

	return router
}
