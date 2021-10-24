package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			temp := math.Pow(float64(i), 2)
			fmt.Printf("%d^2 = %v\n", i, int(temp))
		}
	}
}
