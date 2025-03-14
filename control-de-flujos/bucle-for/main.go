package main

import "fmt"

func main() {
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
