package arreglos_slices

import "fmt"

// Definicmos un slice
// Asignacion de forma directa
var tabla [10]int = [10]int{10, 0, 59, 10}

// MAtriz vector de vectore
var matriz [20][30]int

func MuestroArreglos() {
	// Asignacion directa
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"Pablo", "Juena", "Elena"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	// Recorreolos siempre con for
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[2][24] = 16

	fmt.Println(matriz)

}
