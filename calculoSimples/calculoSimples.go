package main

import (
	"fmt"
)

func main() {
	var cod1 int
	var num1 int
	var valorUnit1 float64
	var cod2 int
	var num2 int
	var valorUnit2 float64

	fmt.Scanf("%v%v%v\n", &cod1, &num1, &valorUnit1)
	fmt.Scanf("%v%v%v\n", &cod2, &num2, &valorUnit2)
	total := (float64(num1) * valorUnit1) + (float64(num2) * valorUnit2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f", total)
}
