package funciones

import "fmt"

// Funciones anonimas
func Calculos() {
	var num3 int = 32
	var num4 int = 243

	calculo := func(num1, num2 int) int {
		return num1 + num2 + num3 + num4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(num1, num2 int) int {
		return num1 * num2 * num3 * num4
	}

	fmt.Println(calculo(10, 25))
}
