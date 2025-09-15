package plays

import (
	"fmt"
	"math/rand"
)

func Play() {
	var numeroIngresado int
	aleatory := rand.Intn(100)
	intentos := 0
	maxIntentos := 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número: (Intentos restantes %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numeroIngresado)

		if numeroIngresado == aleatory {
			fmt.Println("Ganaste!!!!!")
			Replay()
			return
		} else if numeroIngresado > aleatory {
			fmt.Printf("El número es menor a (%d) OJO\n", numeroIngresado)
		} else if numeroIngresado < aleatory {
			fmt.Printf("El número es mayor a (%d) OJO\n", numeroIngresado)
		}
	}

	fmt.Printf("Se acabo el juego el número era: %d\n", aleatory)
	Replay()
}

func Replay() {
	var desicion string
	fmt.Println("¿Quieres jugar de nuevo? (s/n)")
	fmt.Scanln(&desicion)

	switch desicion {
	case "s":
		Play()
	case "n":
		fmt.Println("Muchas gracias por juagar! <3")
	default:
		fmt.Println("Selección erronea =(")
		Replay()
	}
}
