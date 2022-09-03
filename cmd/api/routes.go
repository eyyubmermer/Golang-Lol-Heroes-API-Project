package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/hero/:id", app.getOneHero)
	router.HandlerFunc(http.MethodGet, "/v1/heroes", app.getAllHeroes)
	router.HandlerFunc(http.MethodGet, "/v1/heroes/:genre_id", app.getAllHeroesByGenre)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genre/:id", app.getOneGenre)

	return app.enableCORS(router)
}
