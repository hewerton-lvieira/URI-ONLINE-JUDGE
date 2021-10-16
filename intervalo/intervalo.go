package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Scanln(&a)
	switch {
	case (a >= 0) && (a <= 25):
		fmt.Println("Intervalo [0,25]")
	case (a > 25) && (a <= 50):
		fmt.Println("Intervalo (25,50]")
	case (a > 50) && (a <= 75):
		fmt.Println("Intervalo (50,75]")
	case (a > 75) && (a <= 100):
		fmt.Println("Intervalo (75,100]")
	case (a < 0) || (a > 100):
		fmt.Println("Fora de intervalo")
	}

}
