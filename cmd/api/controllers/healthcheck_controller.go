package controllers

import (
	"fmt"
	"github.com/whoiswentz/greenlight/cmd/api/application"
	"log"
	"net/http"
)

type HealthCheckController struct {
	logger *log.Logger
}

func NewHealthCheckController() HealthCheckController {
	return HealthCheckController{
		logger: application.NewLogger("[HealthCheckController] - "),
	}
}

func (c HealthCheckController) GET(app application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.logger.Printf("%s - %s", r.Method, r.URL)
		fmt.Fprintln(w, "status: available")
		fmt.Fprintf(w, "environment: %s\n", app.Config.Env)
		fmt.Fprintf(w, "version: %s\n", app.Version)
	}
}
