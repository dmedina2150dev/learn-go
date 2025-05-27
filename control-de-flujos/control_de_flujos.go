package controldeflujos

import (
	"fmt"
	"runtime"
)

func ValidateOSByIfElse() {
	os := runtime.GOOS

	if os == "Linux." || os == "darwin" {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	if os1 := runtime.GOOS; os1 == "Linux." || os1 == "darwin" {
		fmt.Println("Esto no es Windows es: ", os1)
	} else {
		fmt.Println("Esto es Windows")
	}
}

func ValidateOSBySwitch() {
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
