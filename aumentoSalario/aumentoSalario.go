package main

import (
	"fmt"
)

func main() {
	var salario, juros float64
	fmt.Scanf("%v\n", &salario)
	var porcento string
	porcento = ("%")
	switch {
	case salario >= 0 && salario <= 400.00:
		juros = 0.15

	case salario > 400 && salario <= 800:
		juros = 0.12

	case salario > 800 && salario <= 1200:
		juros = 0.10

	case salario > 1200 && salario <= 2000:
		juros = 0.07

	case salario > 2000:
		juros = 0.04

	default:
		juros = 0
	}
	novoSalario := (salario * juros) + salario
	valorAumento := novoSalario - salario
	jurosValor := juros * 100
	fmt.Printf("Novo salario: %.2f\n", novoSalario)
	fmt.Printf("Reajuste ganho: %.2f\n", valorAumento)
	fmt.Printf("Em percentual: %d %s\n", int(jurosValor), porcento)

}
