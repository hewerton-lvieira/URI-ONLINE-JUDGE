package main

import (
	"fmt"
)

func valorImposto(a float64) {
	var juros float64
	switch {
	case (a >= 0.00 && a <= 2000.00):
		fmt.Println("Isento")
	case a >= 2000.01 && a <= 3000.00:
		juros = (a - 2000) * 0.08
		fmt.Printf("R$ %.2f\n", juros)
	case a >= 3000.01 && a <= 4500.00:
		if a <= 3000 {
			juros = (a - 2000) * 0.18
		} else {
			juros = (a-3000)*0.18 + (1000 * 0.08)
		}
		fmt.Printf("R$ %.2f\n", juros)
	case a > 4500.01:
		juros = ((a - 4500) * 0.28) + (1000 * 0.08) + (1500 * 0.18)
		fmt.Printf("R$ %.2f\n", juros)
	}
}

func main() {
	var a float64
	fmt.Scanf("%f\n", &a)
	valorImposto(a)
}
