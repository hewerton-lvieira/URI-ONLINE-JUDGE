package main

import (
	"fmt"
)

func main() {
	var diaIni, horaIni, minIni, segIni, diaFin, horaFin, minFin, segFim float64
	fmt.Scanf("Dia %v\n", &diaIni)
	fmt.Scanf("%v : %v : %v\n", &horaIni, &minIni, &segIni)
	fmt.Scanf("Dia %v\n", &diaFin)
	fmt.Scanf("%v : %v : %v\n", &horaFin, &minFin, &segFim)
	quantDiaIni := (diaFin * 86400)
	quantDiaFim := (diaIni * 86400)
	quantHoraIni := (horaFin * 3600)
	quantHoraFim := (horaIni * 3600)
	quantMinIni := (minFin * 60)
	quantMinFim := (minIni * 60)
	quantSegIni := segFim
	quantSegFim := segIni
	totalIni := quantDiaIni + quantHoraIni + quantMinIni + quantSegIni
	totalFim := quantDiaFim + quantHoraFim + quantMinFim + quantSegFim
	tempoSeg := totalFim - totalIni
	if tempoSeg < 0 {
		tempoSeg = tempoSeg * -1
	}
	a := int(tempoSeg) % 86400
	duracaoDias := (int(tempoSeg) - a) / 86400
	b := a % 3600
	duracaoHoras := (a - b) / 3600
	c := b % 60
	duracaoMin := (b - c) / 60
	duracaoSeg := c
	fmt.Println(duracaoDias, "dia(s)")
	fmt.Println(duracaoHoras, "hora(s)")
	if tempoSeg <= 60 {
		duracaoMin = 1
		duracaoSeg = 0
	}
	fmt.Println(duracaoMin, "minuto(s)")
	fmt.Println(duracaoSeg, "segundo(s)")
}
