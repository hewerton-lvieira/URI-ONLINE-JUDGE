package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%v%v%v", &a, &b, &c)
	if a < b+c && b < c+a && c < a+b {
		d := a + b + c
		fmt.Printf("Perimetro = %.1f\n", d)
	} else {
		d := ((a + b) * c) / 2
		fmt.Printf("Area = %.1f\n", d)
	}
}
