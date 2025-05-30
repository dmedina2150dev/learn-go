package ejercicios

import (
	"bufio"
	"fmt"
	"os"
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

func CreateTableByBufio() {
	scanner := bufio.NewScanner(os.Stdin)
	var valorNum int
	var err error

	for {
		fmt.Println("Ingresa un número del entero: ")
		if scanner.Scan() {
			valorNum, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	fmt.Printf("Esta en la tabla del: %d \n", valorNum)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", valorNum, i, valorNum*i)
	}

	fmt.Printf("Fin de la tabla \n")
}

func CreateTableMultiplySave() string {
	scanner := bufio.NewScanner(os.Stdin)
	var valorNum int
	var err error
	var texto string

	for {
		fmt.Println("Ingresa un número del entero: ")
		if scanner.Scan() {
			valorNum, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	texto += fmt.Sprintf("Esta en la tabla del: %d \n", valorNum)

	// TODO: Haremos una paquete para guardar en el archivo
	for i := 1; i <= 10; i++ {
		// TODO: Cambiaremos a SprintF --> Para que nos retorne un string
		texto += fmt.Sprintf("%d x %d = %d \n", valorNum, i, valorNum*i)
	}

	texto += fmt.Sprintf("Fin de la tabla")

	return texto
}
