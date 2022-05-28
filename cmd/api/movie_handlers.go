package main

import (
	"fmt"
	"net/http"
)

func (app *Application) CreateMovie(w http.ResponseWriter, r *http.Request) {
}

func (app *Application) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.Request.GetIntParam(r, "id")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie: %d", id)
}

func (app *Application) DeleteMovie(w http.ResponseWriter, r *http.Request) {

}
