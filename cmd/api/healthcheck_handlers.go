package main

import (
	"net/http"
)

func (app *Application) Check(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.Config.Env,
		"version":     app.Version,
	}

	if err := app.Response.JSON(w, http.StatusOK, data, nil); err != nil {
		app.Logger.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}
