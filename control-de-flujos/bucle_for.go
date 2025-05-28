package controldeflujos

import "fmt"

func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}

	for i := 100; i > 10; i-- {
		if i == 20 {
			break
		}

		fmt.Println(i)
	}
}

func BucleFor() {
	var i int

	for i <= 10 {
		fmt.Println("i:", i)
		i++
	}

	for b := 0; b <= 10; b++ {
		fmt.Println("b:", b)
		if b == 5 {
			break
		}
	}

	for c := 0; c <= 10; c++ {
		if c == 5 {
			continue
		}
		fmt.Println("c:", c)
	}
}
