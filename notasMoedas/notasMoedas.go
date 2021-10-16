package main

import (
	"fmt"
)

func main() {
	var valor float64
	fmt.Scanln(&valor)
	var tirarNotas100 = int(valor) % 100           // TIRAR AS NOTAS DE 100
	notas100 := (int(valor) - tirarNotas100) / 100 // QUANTIDADE DE NOTAS DE 100
	fmt.Println("NOTAS:")
	fmt.Println(notas100, "nota(s) de R$ 100.00")

	tirarNotas50 := tirarNotas100 % 50             // TIRAR AS NOTAS DE 50
	notas50 := (tirarNotas100 - tirarNotas50) / 50 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas50, "nota(s) de R$ 50.00")

	tirarNotas20 := tirarNotas50 % 20
	notas20 := (tirarNotas50 - tirarNotas20) / 20 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas20, "nota(s) de R$ 20.00")

	tirarNotas10 := tirarNotas20 % 10
	notas10 := (tirarNotas20 - tirarNotas10) / 10 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas10, "nota(s) de R$ 10.00")

	tirarNotas5 := tirarNotas10 % 5
	notas5 := (tirarNotas10 - tirarNotas5) / 5 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas5, "nota(s) de R$ 5.00")

	tirarNotas2 := tirarNotas5 % 2
	notas2 := (tirarNotas5 - tirarNotas2) / 2 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas2, "nota(s) de R$ 2.00")

	tirarNotas1 := tirarNotas2 % 1
	notas1 := (tirarNotas2 - tirarNotas1) / 1 // QUANTIDADE DE NOTAS DE 50
	fmt.Println("MOEDAS:")
	fmt.Println(notas1, "moeda(s) de R$ 1.00")
	valor1 := (valor * 100)
	valortemp := int(valor1) % 100
	tirarNotas05 := valortemp % 50
	notas05 := (valortemp - tirarNotas05) / 50 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas05, "moeda(s) de R$ 0.50")
	tirarNotas025 := tirarNotas05 % 25
	notas025 := (tirarNotas05 - tirarNotas025) / 25 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas025, "moeda(s) de R$ 0.25")
	tirarNotas010 := tirarNotas025 % 10
	notas010 := (tirarNotas025 - tirarNotas010) / 10 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas010, "moeda(s) de R$ 0.10")
	tirarNotas005 := tirarNotas010 % 5
	notas005 := (tirarNotas010 - tirarNotas005) / 5 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas005, "moeda(s) de R$ 0.05")
	tirarNotas001 := tirarNotas005 % 1
	notas001 := (tirarNotas005 - tirarNotas001) / 1 // QUANTIDADE DE NOTAS DE 50
	fmt.Println(notas001, "moeda(s) de R$ 0.01")
}
