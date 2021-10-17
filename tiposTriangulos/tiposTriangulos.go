package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%v%v%v\n", &a, &b, &c)
	s1 := []float64{a, b, c}
	sort.Float64s(s1)
	var s2 []float64
	for i := len(s1) - 1; i >= 0; i-- {
		s2 = append(s2, s1[i])
	}
	aa := s2[0]
	bb := s2[1]
	cc := s2[2]
	if aa >= bb+cc {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if math.Pow(aa, 2) == math.Pow(bb, 2)+math.Pow(cc, 2) {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if math.Pow(aa, 2) > (math.Pow(bb, 2) + math.Pow(cc, 2)) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if math.Pow(aa, 2) < math.Pow(bb, 2)+math.Pow(cc, 2) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if aa == bb && bb == cc && cc == aa {
			fmt.Println("TRIANGULO EQUILATERO")
		}
		if (aa == bb && cc != bb) || (cc == aa && bb != cc) || (cc == bb && aa != bb) {
			fmt.Println("TRIANGULO ISOSCELES")
		}
	}

}
