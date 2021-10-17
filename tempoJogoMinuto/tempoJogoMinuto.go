package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, minuto, hora int
	a, b, c, d = 0, 0, 0, 0
	fmt.Scanf("%d%d%d%d\n", &a, &b, &c, &d)
	if a > c {
		hora = (24 - a) + c
	} else if a < c {
		hora = c - a
	} else {
		hora = 24
	}
	if b > d {
		minuto = (60 - b) + d
	} else if d > b {
		minuto = d - b
	} else {
		minuto = 0
	}
	if d < b {
		hora = hora - 1
	}

	if hora == 0 && minuto == 0 {
		hora = 24
	}
	if hora == 24 && minuto != 0 {
		hora = 0
	}
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", hora, minuto)
}
