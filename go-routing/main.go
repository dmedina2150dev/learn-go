package main

import (
	"fmt"
	"time"
)

func hola() {
	fmt.Println("Hola esto es una gorutina")
}

func main() {
	go hola()
	time.Sleep(1 * time.Second)
	fmt.Println("Hola esto es el main")
}
