package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var d int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	diferenca := (a * b) - (c * d)
	fmt.Println("DIFERENCA =", diferenca)
}
