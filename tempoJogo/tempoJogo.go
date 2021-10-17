package main

import (
	"fmt"
)

func main() {
	var a, b, contDias int
	fmt.Scanf("%v%v\n", &a, &b)
	if a < b {
		contDias = b - a
	} else if a > b {
		contDias = (24 - a) + b
	} else {
		contDias = 24
	}
	fmt.Printf("O JOGO DUROU %v HORA(S)\n", contDias)
}
