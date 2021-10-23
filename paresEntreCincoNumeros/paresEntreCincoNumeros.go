package main

import (
	"fmt"
)

func main() {
	var numeros [5]int
	fmt.Scanln(&numeros[0])
	fmt.Scanln(&numeros[1])
	fmt.Scanln(&numeros[2])
	fmt.Scanln(&numeros[3])
	fmt.Scanln(&numeros[4])
	var cont int
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			cont++
		}
	}
	fmt.Println(cont, "valores pares")
}
