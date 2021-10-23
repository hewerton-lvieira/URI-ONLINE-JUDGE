package main

import (
	"fmt"
)

func main() {
	var numeros [5]int
	fmt.Scanln(&numeros[0])
	fmt.Scanln(&numeros[1])
	fmt.Scanln(&numeros[2])
	fmt.Scanln(&numeros[3])
	fmt.Scanln(&numeros[4])
	var contPar, contImpar, contPos, contNeg int
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			contPar++
		} else {
			contImpar++
		}
		if numeros[i] > 0 {
			contPos++
		} else if numeros[i] < 0 {
			contNeg++
		}
	}
	fmt.Printf("%v valor(es) par(es)\n%v valor(es) impar(es)\n%v valor(es) positivo(s)\n%v valor(es) negativo(s)\n", contPar, contImpar, contPos, contNeg)

}
