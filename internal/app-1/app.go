package app1

import "rsc.io/quote"

type App1Model struct {
	Saludo    string
	SaludoLib string
}

func Saludos() *App1Model {
	saludo := "Hola Mundo!!!"
	saludoLib := quote.Hello()

	return &App1Model{
		Saludo:    saludo,
		SaludoLib: saludoLib,
	}
}
