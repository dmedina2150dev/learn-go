package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dmedina2150dev/learn-go/ejercicios"
)

var fileName string = "./files/tablas/tablas.txt"

func SaveTabla() {
	var texto string = ejercicios.CreateTableMultiplySave()
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo: " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.CreateTableMultiplySave()

	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(fileName, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error al append: " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto + "\n\n")

	if err != nil {
		fmt.Println("Error al escribir: " + err.Error())
		return false
	}

	arch.Close()

	return true
}

func ReadFile() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()

		if strings.Contains(registro, "tabla") {
			fmt.Println(registro)
			continue
		}

		fmt.Println("> " + registro)
	}
	archivo.Close()
}
