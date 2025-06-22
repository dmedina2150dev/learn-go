package mapas

import "fmt"

func MostrarMapas() {

	// Creamos un mapa dinamico con make
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["Venezuela"] = "Cararas"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	// Crear mapa con asignación directa
	campeonato := map[string]int{
		"Real Madrid":        49,
		"Barcelona":          30,
		"Atletico de Madrid": 29,
		"Real Sociedad":      28,
	}

	fmt.Println(campeonato["Real Madrid"])

	// Iteramos el mapa
	for equipo, puntos := range campeonato {
		fmt.Printf("%s, tiene %d puntos \n", equipo, puntos)
	}

	// Eliminar elemento del mapa
	delete(campeonato, "Real Sociedad")
	fmt.Println(campeonato)

	// Buscar un elemento del mapa
	/**
	* LO que devuelve al buscar un elemento en un mapa, son dos variables
	* el primero es el @value del key del elemento
	* el segundo es un valor booleano que indica si el elemento existe o no en el mapa
	 */

	puntos, existe := campeonato["Real Madrid"]
	puntos2, existe2 := campeonato["Juventud"]

	fmt.Printf("El puntaje encontrado es %d, y el equipo existe = %t \n", puntos, existe)
	fmt.Printf("El puntaje encontrado es %d, y el equipo existe = %t \n", puntos2, existe2)

}
