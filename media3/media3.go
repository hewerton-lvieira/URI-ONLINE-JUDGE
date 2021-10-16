package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, n4 float64
	fmt.Scanf("%f%f%f%f\n", &n1, &n2, &n3, &n4)
	media := ((n1 * 2) + (n2 * 3) + (n3 * 4) + (n4 * 1)) / 10
	notaExame := 0.0
	if media >= 5 && media <= 6.9 {
		fmt.Scanf("%v\n", &notaExame)
	}

	fmt.Printf("Media: %.1f\n", media)
	if media >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if media < 5 {
		fmt.Println("Aluno reprovado.")
	} else if media >= 5 && media <= 6.9 {
		fmt.Println("Aluno em exame.")
		fmt.Printf("Nota do exame: %.1f\n", notaExame)
		var novaMedia float64
		novaMedia = (notaExame + media) / 2
		if novaMedia >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", novaMedia)
	}
}
