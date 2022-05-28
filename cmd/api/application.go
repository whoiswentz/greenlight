package main

import (
	"log"
)

type Application struct {
	Version  string
	Config   Config
	Logger   *log.Logger
	Request  *request
	Response *response
}

func NewApplication(version string, config Config, logger *log.Logger) Application {
	return Application{
		Version:  version,
		Config:   config,
		Logger:   logger,
		Request:  &request{},
		Response: &response{},
	}
}
