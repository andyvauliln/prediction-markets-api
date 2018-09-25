package main

import (
	"fmt"
	"testing"

	"github.com/andyvauliln/prediction-markets-api/app"
	"github.com/andyvauliln/prediction-markets-api/e2e"
	"github.com/andyvauliln/prediction-markets-api/errors"
)

func TestApp(t *testing.T) {
	// load application configurations
	if err := app.LoadConfig("./config", ""); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	app.Config.DBName = "proofdextest"
	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	e2e.Init(t)
}
