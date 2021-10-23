package main

import (
	"fmt"
)

func main() {
	var x, cont int
	fmt.Scanln(&x)
	for i := x; i <= x+50; i++ {
		if i%2 != 0 && cont < 6 {
			fmt.Println(i)
			cont++
		}
	}
}
