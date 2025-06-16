package arreglos_slices

import "fmt"

// Un Slice Es una arreglo dinamico
var tablaS []int = []int{2, 5, 4}

var arreglo [10]int = [10]int{5, 78, 90, 54, 56, 789, 21}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // Slice creado desde un vector desde la pocision 3
	porcion2 := arreglo[:5]  // Slice creado desde un vector desde la pocision posicion 0 hasta la 5
	porcion3 := arreglo[1:2] // Slice creado desde un vector desde la pocision posicion 1 hasta la 2

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	// make -->
	/**
	* make params
	* @param slice
	* @param capacidad inicial
	* @param capacidad de aumentar
	 */
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0) // O se puede hacer de esta forma

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
