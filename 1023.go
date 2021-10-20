package main

import (
	"fmt"
)

func main() {
	var A float64
	fmt.Scanf("%f", &A)
	if A >= 0 && A <= 25 {
		fmt.Printf("Intervalo [0,25]\n")
	} else if A > 25 && A <= 50 {
		fmt.Printf("Intervalo (25,50]\n")
	} else if A > 50 && A <= 75 {
		fmt.Printf("Intervalo (50,75]\n")
	} else if A > 75 && A <= 100 {
		fmt.Printf("Intervalo (75,100]\n")
	} else {
		fmt.Printf("Fora de intervalo\n")
	}
}
