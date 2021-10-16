package main

import (
	"fmt"
)

func main() {
	var valorTotal int
	fmt.Scanln(&valorTotal)
	notas := valorTotal % 100
	quantNotas := (valorTotal - notas) / 100
	fmt.Println(valorTotal)
	fmt.Println(quantNotas, "nota(s) de R$ 100,00")
	valorTotal = notas
	notas = valorTotal % 50
	quantNotas = (valorTotal - notas) / 50
	fmt.Println(quantNotas, "nota(s) de R$ 50,00")
	valorTotal = notas
	notas = valorTotal % 20
	quantNotas = (valorTotal - notas) / 20
	fmt.Println(quantNotas, "nota(s) de R$ 20,00")
	valorTotal = notas
	notas = valorTotal % 10
	quantNotas = (valorTotal - notas) / 10
	fmt.Println(quantNotas, "nota(s) de R$ 10,00")
	valorTotal = notas
	notas = valorTotal % 5
	quantNotas = (valorTotal - notas) / 5
	fmt.Println(quantNotas, "nota(s) de R$ 5,00")
	valorTotal = notas
	notas = valorTotal % 2
	quantNotas = (valorTotal - notas) / 2
	fmt.Println(quantNotas, "nota(s) de R$ 2,00")
	valorTotal = notas
	notas = valorTotal % 1
	quantNotas = (valorTotal - notas) / 1
	fmt.Println(quantNotas, "nota(s) de R$ 1,00")
}
