package main

import (
	"fmt"
)

func main() {
	var tempoViagem int
	var velocidadeMedia int

	fmt.Scanln(&tempoViagem)
	fmt.Scanln(&velocidadeMedia)

	var consumoTotal float64 = (1 * (float64(tempoViagem) * float64(velocidadeMedia))) / 12
	fmt.Printf("%.3f\n", consumoTotal)

}
