package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%v\n", &a)
	var b int
	fmt.Scanln(&b)
	soma := a + b
	fmt.Println("SOMA =", soma)
}
