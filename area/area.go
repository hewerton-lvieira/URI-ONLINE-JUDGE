package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	pi := 3.14159
	fmt.Scanf("%v%v%v\n", &a, &b, &c)
	tri := (a * c) / 2
	ci := math.Pow(c, 2) * pi
	tra := ((a + b) * c) / 2
	qua := math.Pow(b, 2)
	ret := a * b

	fmt.Printf("TRIANGULO: %.3f\n", tri)
	fmt.Printf("CIRCULO: %.3f\n", ci)
	fmt.Printf("TRAPEZIO: %.3f\n", tra)
	fmt.Printf("QUADRADO: %.3f\n", qua)
	fmt.Printf("RETANGULO: %.3f\n", ret)
}
