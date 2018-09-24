package main

import (
	"fmt"
	"testing"

	"github.com/andyvaulin/prediction-markets/app"
	"github.com/andyvaulin/prediction-markets/e2e"
	"github.com/andyvaulin/prediction-markets/errors"
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
