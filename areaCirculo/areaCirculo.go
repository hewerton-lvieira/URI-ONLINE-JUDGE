package main

import (
	"fmt"
	"math"
)

func main() {
	var area float64
	var n float64
	var raio float64
	n = 3.14159
	fmt.Scanln(&raio)

	area = n * math.Pow(raio, 2)
	fmt.Printf("A=%.4f\n", area)

}
