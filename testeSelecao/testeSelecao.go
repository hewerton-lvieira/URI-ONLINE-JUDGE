package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var d int

	fmt.Scanf("%v%v%v%v", &a, &b, &c, &d)

	if b > c && d > a && (c+d) > (a+b) && c >= 0 && d >= 0 && a%2 == 0 {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}
