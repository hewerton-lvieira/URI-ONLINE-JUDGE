package main

import (
	"fmt"
)

func main() {
	var distancia int
	var combustivel float64
	fmt.Scanln(&distancia)
	fmt.Scanln(&combustivel)
	resultado := float64(distancia) / combustivel
	fmt.Printf("%.3f km/l\n", resultado)
}
