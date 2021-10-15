package main

import (
	"fmt"
)

func main() {
	var hora int
	var numero int
	var receber float64

	fmt.Scanln(&numero)
	fmt.Scanln(&hora)
	fmt.Scanln(&receber)

	salario := float64(hora) * receber
	fmt.Println("NUMBER =", numero)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
