package arreglos_slices

import "fmt"

func UnderArrays() {
	// 1. Declarar matriz [cantidad_de_elementos]tipo_de_los_Elementos
	var nameMatriz [5]int

	// 2. Modificar valores de la matriz (Se puede usar el indice de cada elemento)
	nameMatriz[1] = 10
	nameMatriz[0] = 1
	fmt.Println(nameMatriz)

	// 3. Declarar e iniciarlizar la matriz
	var matriz2 = [10]int{0, 12, 0, 9, 2} // El resto se inicializa en 0
	fmt.Println(matriz2)

	// 4. Declarar e inicializar la matriz sin conocer la cantidad de elementos que contiene
	var matriz = [...]int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(matriz)

	// 5. Iterar los elementos de una matriz
	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	for index, value := range matriz {
		fmt.Printf("Indice %d, valor: %d\n", index, value)
	}

	// 6. Matriz Bidicemcional
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(m)

}
