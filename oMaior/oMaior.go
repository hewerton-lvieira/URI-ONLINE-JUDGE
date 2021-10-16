package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Scanf("%v%v%v\n", &a, &b, &c)
	af := float64(a)
	bf := float64(b)
	cf := float64(c)
	maiorAB := math.Max(af, bf)
	resultado := math.Max(maiorAB, cf)
	fmt.Printf("%v eh o maior\n", resultado)
}
