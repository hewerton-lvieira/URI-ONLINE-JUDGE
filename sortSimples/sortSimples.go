package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%v%v%v", &a, &b, &c)
	s1 := []int{a, b, c}
	sort.Ints(s1)
	for j := 0; j < len(s1); j++ {
		fmt.Println(s1[j])
	}
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
