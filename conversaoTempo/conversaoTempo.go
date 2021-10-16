package main

import (
	"fmt"
)

func main() {
	var totalsegundos int
	fmt.Scanln(&totalsegundos)
	segundos := totalsegundos % 60
	conversaoMinutos := (totalsegundos - segundos) / 60
	minutos := conversaoMinutos % 60
	conversaohoras := (conversaoMinutos - minutos) / 60

	fmt.Printf("%v:%v:%v", conversaohoras, minutos, segundos)
}
