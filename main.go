package main

import (
	"fmt"

	hello "github.com/dmedina2150dev/learn-go/hola-mundo"
	"github.com/dmedina2150dev/learn-go/variables"
	"github.com/dmedina2150dev/learn-go/ejercicios"
)

func main() {
	fmt.Println("Hola Mundo!!!")
	hello.SayHello()
	variables.MuestroEnteros()
	variables.RestoVariables()

	state, text := variables.ConvertToText(1950)

	fmt.Println(state)
	fmt.Println(text)

	fmt.Print("\n\n")

	var dato string
	fmt.Print("Ingrea un número entero:")
	fmt.Scanln(&dato)

	value, msg := ejercicios.CalculateAndTransform(dato)

	fmt.Printf("valor: %d, y %s \n", value, msg)
}
