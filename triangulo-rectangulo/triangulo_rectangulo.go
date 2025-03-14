package triangulo_rectangulo

import (
	"fmt"
	"math"
)

func TrianguloRectangulo() {
	// Definicion de variables necesarias
	var lado1, lado2 float64
	const presicion = 2

	// Entrada de datos
	fmt.Print("Ingrese el valor del lado 1: ")
	fmt.Scanln(&lado1) // El signo & se utiliza para obtener la direccion de memoria de la variable
	fmt.Print("Ingrese el valor del lado 2: ")
	fmt.Scanln(&lado2) // El signo & se utiliza para obtener la direccion de memoria de la variable

	// Proceso

	area := (lado1 * lado2) / 2

	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))

	perimetro := lado1 + lado2 + hipotenusa

	// Salida
	fmt.Printf("\nEl area del triangulo rectangulo es: %.2f \n", area)                      // Una forma de acortar los decimales que se imprimen
	fmt.Printf("\nEl perimetro del triangulo rectangulo es: %.*f \n", presicion, perimetro) // Otra forma de acortar los decimales que se imprimen
}
