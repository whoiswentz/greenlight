package application

import (
	"github.com/whoiswentz/greenlight/cmd/api/configuration"
)

type Application struct {
	Version string
	Config  configuration.Config
	Request request
}

func NewApplication(version string, config configuration.Config) Application {
	return Application{
		Version: version,
		Config:  config,
		Request: request{},
	}
}
