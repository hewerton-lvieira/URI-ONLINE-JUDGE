package main

import (
	"fmt"
)

func main() {
	var n, contIntervalo, x, contForaInter int
	fmt.Scanf("%v\n", &n)
	var s1 []int
	for i := 1; i <= n; i++ {
		fmt.Scanln(&x)
		s1 = append(s1, x)
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] <= 20 && s1[i] >= 10 {
			contIntervalo++

		} else {
			contForaInter++
		}
	}
	fmt.Println(contIntervalo, "in")
	fmt.Println(contForaInter, "out")
}
