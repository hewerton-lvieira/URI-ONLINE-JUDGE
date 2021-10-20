package main

import (
	"fmt"
)

func escolha(a int) string {
	var mes [12]string = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	b := a - 1
	return mes[b]

}
func main() {
	var a int
	fmt.Scanf("%v\n", &a)
	b := escolha(a)
	fmt.Println(b)
}
