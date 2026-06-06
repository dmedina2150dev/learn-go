package ejercicios

import "strconv"

func CalculateAndTransform(value string) (int, string) {
	convert, err := strconv.Atoi(value)

	if err != nil {
		return 0, "Ups! Ocurrio un error: " + err.Error()
	}

	if convert > 100 {
		return convert, "Es mayor a 100"
	} else {
		return convert, "Es menor a 100"
	}
}
