package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f, cont float64
	fmt.Scanf("%v\n%v\n%v\n%v\n%v\n%v\n", &a, &b, &c, &d, &e, &f)
	var s1 []float64
	s1 = []float64{a, b, c, d, e, f}
	for i := 0; i != len(s1); i++ {
		switch {
		case s1[i]*1 >= 0:
			cont++
		default:

		}
	}
	fmt.Printf("%v valores positivos\n", cont)
}
