package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	for i := 0; i <= x; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
