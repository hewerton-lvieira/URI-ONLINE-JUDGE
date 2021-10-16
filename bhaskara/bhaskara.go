package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Scanf("%v%v%v", &a, &b, &c)

	delta := math.Pow(b, 2) - 4*a*c
	d := math.Sqrt(delta)
	x1 := (-b + d) / (2 * a)
	x2 := (-b - d) / (2 * a)
	if delta == 0 || math.IsNaN(x1) || math.IsNaN(x2) {
		fmt.Println("Impossivel calcular")
	} else {
		fmt.Printf("R1 = %.5f\n", x1)
		fmt.Printf("R2 = %.5f\n", x2)
	}
}
