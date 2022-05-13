package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/whoiswentz/greenlight/cmd/api/application"
	"github.com/whoiswentz/greenlight/cmd/api/configuration"
	"github.com/whoiswentz/greenlight/cmd/api/controllers"
	"net/http"
	"time"
)

const version = "1.0.0"

func main() {
	config := configuration.NewConfig()

	flag.IntVar(&config.Port, "port", 4000, "API server port")
	flag.StringVar(&config.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := application.NewLogger("[MAIN] - ")

	app := application.NewApplication(version, config)

	healthCheckHandler := controllers.NewHealthCheckController()
	movieHandler := controllers.NewMovieHandler(&app)

	mux := httprouter.New()
	mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthCheckHandler.GET(app))
	mux.HandlerFunc(http.MethodPost, "/v1/movies", movieHandler.CreateMovie)
	mux.HandlerFunc(http.MethodGet, "/v1/movies/:id", movieHandler.GetMovieByID)

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
