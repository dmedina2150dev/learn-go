package controldeflujos

import (
	"fmt"
	"runtime"
	"time"
)

func GameGuessNumber() {
	t := time.Now() // Obtiene la hora actual
	hora := t.Hour()

	fmt.Println(t)                               // Imprime el tiempo actual
	fmt.Println(t.Year())                        // Imprime el año actual
	fmt.Println(t.Month())                       // Imprime el mes actual
	fmt.Println(t.Day())                         // Imprime el día actual
	fmt.Println(t.Hour())                        // Imprime la hora actual
	fmt.Println(t.Minute())                      // Imprime el minuto actual
	fmt.Println(t.Second())                      // Imprime el segundo actual
	fmt.Println(t.Nanosecond())                  // Imprime el nanosegundo actual
	fmt.Println(t.Location())                    // Imprime la ubicación actual
	fmt.Println(t.Weekday())                     // Imprime el día de la semana actual
	fmt.Println(t.YearDay())                     // Imprime el día del año actual
	fmt.Println(t.Format("02/01/2006 15:04:05")) // Imprime la fecha y hora en formato personalizado

	if hora < 12 {
		fmt.Println("Buenos días - Es de mañana")
	} else if hora < 17 {
		fmt.Println("Buenas tardes - Es de tarde")
	} else {
		fmt.Println("Buenas noches - Es de noche")
	}

	// OTRA FORMA DE HACERLO

	if tiempo := time.Now(); tiempo.Hour() < 12 {
		fmt.Println("Buenos días - Es de mañana")
	} else if tiempo.Hour() < 17 {
		fmt.Println("Buenas tardes - Es de tarde")
	} else {
		fmt.Println("Buenas noches - Es de noche")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Buenos días - Es de mañana --> Con switch")
	case t.Hour() < 17:
		fmt.Println("Buenas tardes - Es de tarde --> Con switch")
	default:
		fmt.Println("Buenas noches - Es de noche --> Con switch")
	}

	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Windows.")
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// OTRA FORMA DE HACERLO

	switch system := runtime.GOOS; system {
	case "windows":
		fmt.Println("Windows.")
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", system)
	}
}
