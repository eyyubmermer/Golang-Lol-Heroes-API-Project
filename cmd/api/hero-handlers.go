package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneHero(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	hero, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, hero, "hero")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getAllHeroes(w http.ResponseWriter, r *http.Request) {

	heroes, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, heroes, "heroes")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllHeroesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	heroID, err := strconv.Atoi(params.ByName("hero_id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	heroes, err := app.models.DB.All(heroID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, heroes, "hero")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
// }

// func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {
// }

// func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
// }

// func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {
// }

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {

	genres, err := app.models.DB.AllGenres()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getOneGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	genre, err := app.models.DB.GetGenre(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, genre, "genre")
	if err != nil {
		app.logger.Println(err)
	}
}
