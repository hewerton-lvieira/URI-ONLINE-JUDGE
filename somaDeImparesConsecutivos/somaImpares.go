package main

import (
	"fmt"
)

func main() {
	var x, y, i, somatorio int
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if x < y {
		for i = x + 1; i < y; i++ {
			if i%2 != 0 {
				somatorio = somatorio + i

			}
		}
	} else if y < x || y == x {
		for i := y + 1; i < x; i++ {
			if i%2 != 0 {
				somatorio = somatorio + i
			}
		}
	}
	fmt.Println(somatorio)
}
