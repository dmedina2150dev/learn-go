package arreglos_slices

import "fmt"

func UnderSlices() {
	//* 1. Declarar slice
	var a []int
	fmt.Println(a)

	//* 2. Añadir elementos a un slice
	a = append(a, 1)
	fmt.Println(a)

	//* 3. Declarar y asignar valores a un slice
	diasSemana := []string{
		"Domingo",
		"Lunes",
		"Martes",
		"Miercoles",
		"Jueves",
		"Viernes",
		"Sabado",
	}
	fmt.Println(diasSemana)

	//* 4. A partir de un slice, podemos crear otro slice ---> Los (:) representa desde antes de los : hasta despues de los dos punto
	diaSlice := diasSemana[0:4]
	fmt.Println(diaSlice)

	diaSlice = append(diaSlice, "Jueves", "Viernes")
	fmt.Println(diaSlice)

	//* 5. Conocer la longitud y la capacidad  que puede almacenar un slice
	fmt.Println(len(diaSlice)) // --> Longitud
	fmt.Println(cap(diaSlice)) // --> Capacidad

	//* 6 Eliminar elementos de un slice
	diaSlice = append(diaSlice[:2], diaSlice[3:]...)
	fmt.Println(diaSlice)

	//* 7. Uso de make para crear slices
	nombres := make([]string, 5)
	fmt.Println(nombres)

	//* 8 Copiar elementos de un slice a otro
	slices1 := []int{1, 2, 4, 6, 8}
	slices2 := make([]int, 5)

	copy(slices1, slices2)

	fmt.Println(slices1)
	fmt.Println(slices2)

	slices3 := []int{1, 2, 4, 6, 8}
	slices4 := make([]int, 5)
	copy(slices4, slices3)

	fmt.Println(slices3)
	fmt.Println(slices4)

}
