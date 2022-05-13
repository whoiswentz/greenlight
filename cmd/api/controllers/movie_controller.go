package controllers

import (
	"fmt"
	"github.com/whoiswentz/greenlight/cmd/api/application"
	"log"
	"net/http"
)

type MovieHandler struct {
	logger *log.Logger
	app    *application.Application
}

func NewMovieHandler(app *application.Application) *MovieHandler {
	return &MovieHandler{
		logger: application.NewLogger("[MovieController - ]"),
		app:    app,
	}
}

func (c *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
}

func (c *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id, err := c.app.Request.GetID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie: %d", id)
}

func (c *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {

}
