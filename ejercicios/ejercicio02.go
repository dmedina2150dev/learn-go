package ejercicios

import (
	"fmt"
	"strconv"
)

var valor string

func CreateTable() {
	var valorNum int
	var err error

	for {
		fmt.Println("Ingresa un número del entero: ")
		fmt.Scanln(&valor)

		valorNum, err = strconv.Atoi(valor)

		if err != nil {
			continue
		}

		break
	}

	fmt.Printf("Esta en la tabla del: %d \n", valorNum)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", valorNum, i, valorNum*i)
	}

	fmt.Printf("Fin de la tabla \n")

}
