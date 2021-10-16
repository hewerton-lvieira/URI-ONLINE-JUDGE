package main

import (
	"fmt"
)

func main() {
	tempoPadrao := 60 //1h igual 60 min
	distPadrao := 30  //30km por hora
	var distFornecida int
	fmt.Scanln(&distFornecida) //distancia informada
	curiosidade := (tempoPadrao * distFornecida) / distPadrao
	fmt.Println(curiosidade, "minutos")

}
