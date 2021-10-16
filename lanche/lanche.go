package main

import (
	"fmt"
)

func main() {
	var infCodigo int
	var quantProduto int
	//codArray := [5]int{1, 2, 3, 4, 5}
	//produtoArray := [5]string{"Cachorro Quente", "X-Salada", "X-Bacon", "Torrada simples", "Refrigerante"}
	precoArrey := [5]float64{4, 4.5, 5, 2, 1.5}
	fmt.Scanf("%v%v", &infCodigo, &quantProduto)
	var total float64
	switch {
	case infCodigo == 1:
		total = float64(quantProduto) * precoArrey[infCodigo-1]
		fmt.Printf("Total: R$ %.2f\n", total)
	case infCodigo == 2:
		total = float64(quantProduto) * precoArrey[infCodigo-1]
		fmt.Printf("Total: R$ %.2f\n", total)
	case infCodigo == 3:
		var numero int = infCodigo - 1
		total = float64(quantProduto) * precoArrey[numero]
		fmt.Printf("Total: R$ %.2f\n", total)
	case infCodigo == 4:
		total = float64(quantProduto) * precoArrey[infCodigo-1]
		fmt.Printf("Total: R$ %.2f\n", total)
	case infCodigo == 5:
		total = float64(quantProduto) * precoArrey[infCodigo-1]
		fmt.Printf("Total: R$ %.2f\n", total)
	}

}
