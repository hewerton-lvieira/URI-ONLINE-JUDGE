package main

import "fmt"

func main() {
	var idadeDias int
	fmt.Scanln(&idadeDias)           // VALOR A SER TRATADO
	dias1part := idadeDias % 365     //EXCLUINDO OS ANOS E SOBRANDO MESES E DIAS
	dias := dias1part % 30           //EXCLUINDO OS MESES E SOBRANDO OS DIAS
	meses := (dias1part - dias) / 30 //EXCLUINDO OS DIAS E SOBRANDO OS MESES
	anos := (idadeDias - (meses * 30) - dias) / 365
	fmt.Println(anos, "ano(s)")
	fmt.Println(meses, "mes(es)")
	fmt.Println(dias, "dia(s)")
}
