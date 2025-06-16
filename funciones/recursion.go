package funciones

import "fmt"

// Funcion que se llama a si misma varias veces

func Exponencia(valor int) {
	if valor > 10000000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)
}
