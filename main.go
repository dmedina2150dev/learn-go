package main

import (
	"fmt"

	"github.com/dmedina2150dev/learn-go/library/book"
	// "github.com/dmedina2150dev/learn-go/server"
	// "github.com/dmedina2150dev/learn-go/users"
	// "github.com/dmedina2150dev/learn-go/mapas"
	// "github.com/dmedina2150dev/learn-go/files"
	// "github.com/dmedina2150dev/learn-go/funciones"
	// "github.com/dmedina2150dev/learn-go/ejercicios"
	// "github.com/dmedina2150dev/learn-go/arreglos_slices"
	// "github.com/dmedina2150dev/learn-go/plays"
	// "github.com/dmedina2150dev/learn-go/variables"
	// "github.com/dmedina2150dev/learn-go/funciones"
	// hello "github.com/dmedina2150dev/learn-go/hola-mundo"
	// controldeflujos "github.com/dmedina2150dev/learn-go/control-de-flujos"
)

func main() {
	fmt.Println("Iniciando GO!")
	fmt.Print("\n\n")
	// hello.SayHello()
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// state, text := variables.ConvertToText(1950)
	// fmt.Println(state)
	// fmt.Println(text)
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
	myBook := book.NewBook("Tu mama fue mia", "Tu mama", 200)

	myBook.PrintInfo()
}
