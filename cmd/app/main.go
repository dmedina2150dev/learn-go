package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "app startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Iniciando GO!")

	resources := InitResources()

	fmt.Println()
	fmt.Println("Saludos -> Primeros pasos --->")
	fmt.Println(resources.hello.Saludo)
	fmt.Println(resources.hello.SaludoLib)
	fmt.Println()

	fmt.Println()
	fmt.Println("Impresión Variables --->")
	resources.variables.VarsNumeric()
	resources.variables.RestVars()
	state, text := resources.variables.Convert(1950)
	fmt.Println(state)
	fmt.Println(text)
	fmt.Println()

	// var dato string
	// fmt.Print("Ingrea un número entero:")
	// fmt.Scanln(&dato)
	// value, msg := ejercicios.CalculateAndTransform(dato)
	// fmt.Printf("valor: %d, y %s \n", value, msg)
	// controldeflujos.Iterar()
	// ejercicios.CreateTable()
	// fmt.Println(ejercicios.CreateTableMultiplySave())
	// files.SaveTabla()
	// files.SumaTabla()
	// files.ReadFile()
	// plays.Play()

	// Clousure
	// funciones.Calculos()
	// funciones.CallClousure()
	// funciones.Exponencia(2)
	// arreglos_slices.MuestroArreglos()
	// arreglos_slices.MuestroSlices()
	// arreglos_slices.Capacidad()
	// arreglos_slices.UnderSlices()
	// mapas.MostrarMapas()
	// mapas.UndeMap()
	// users.AltaUsuario()

	// server.InitServer()

	// Aqui estaria mal si las propiedades del struc son privadas
	// myBook2 := book.Book{
	// 	title:  "Tu mama fue mia",
	// 	author: "Tu mama",
	// 	pages:  200,
	// }
	// myBook2 := book.Book{
	// 	"Tu mama fue mia",
	// 	"Tu mama",
	// 	200,
	// }
	// myBook := book.NewBook("Tu mama fue mia", "Tu mama", 200)

	// myBook.PrintInfo()

	return nil
}
