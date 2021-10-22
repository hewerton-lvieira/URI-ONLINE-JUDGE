package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f float64
	fmt.Scanf("%v\n%v\n%v\n%v\n%v\n%v\n", &a, &b, &c, &d, &e, &f)
	var a1 [6]float64 = [6]float64{a, b, c, d, e, f}
	var contPos int
	var somatorio float64
	for i := 0; i <= (len(a1) - 1); i++ {
		if a1[i] >= 0 {
			contPos++
			somatorio = somatorio + a1[i]
		}
	}
	media := somatorio / float64(contPos)
	fmt.Println(contPos, "valores positivos")
	fmt.Printf("%.1f\n", media)
}
