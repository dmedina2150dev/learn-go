package main

import (
	app1 "github.com/dmedina2150dev/learn-go/internal/app-1"
	"github.com/dmedina2150dev/learn-go/internal/platform/variables"
)

type Resources struct {
	hello     *app1.App1Model
	variables *variables.VariablesModel
}

func InitResources() *Resources {
	// Primera app saludando
	appSaludos := app1.Saludos()

	// Variables
	funcionesVariables := variables.NewClient()

	return &Resources{
		hello:     appSaludos,
		variables: funcionesVariables,
	}
}
