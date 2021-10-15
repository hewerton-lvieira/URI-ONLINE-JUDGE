package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	prod := a * b
	fmt.Println("PROD =", prod)
}
