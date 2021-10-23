package main

import (
	"fmt"
)

func main() {
	var x, y, somatorio int
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if x < y {
		for i := x; i < y; i++ {
			if i%2 != 0 {
				somatorio = somatorio + i

			}
		}
	} else if y < x || y == x {
		for i := y; i < x; i++ {
			if i%2 != 0 {
				fmt.Println("porra")
				somatorio = somatorio + i
			}
		}
	} else {
		somatorio = 80
	}
	fmt.Println(somatorio)
}
