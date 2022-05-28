package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

const version = "1.0.0"

func main() {
	config := NewConfig()

	flag.IntVar(&config.Port, "port", 4000, "API server port")
	flag.StringVar(&config.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := NewLogger("[MAIN] - ")

	app := NewApplication(version, config, logger)

	mux := httprouter.New()
	mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.Check)
	mux.HandlerFunc(http.MethodPost, "/v1/movies", app.CreateMovie)
	mux.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.GetMovieByID)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", config.Env, srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
