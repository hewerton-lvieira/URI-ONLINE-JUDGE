package main

import (
	"fmt"
)

func animal(a, b, c string) string {
	var d string
	if a == "vertebrado" {
		if b == "ave" {
			if c == "carnivoro" {
				d = "aguia"
			} else if c == "onivoro" {
				d = "pomba"
			}
		} else if b == "mamifero" {
			if c == "onivoro" {
				d = "homem"
			} else if c == "herbivoro" {
				d = "vaca"
			}
		}
	} else if a == "invertebrado" {
		if b == "inseto" {
			if c == "hematofago" {
				d = "pulga"
			} else if c == "herbivoro" {
				d = "lagarta"
			}
		} else if b == "anelideo" {
			if c == "hematofago" {
				d = "sanguessuga"
			} else if c == "onivoro" {
				d = "minhoca"
			}

		}

	} else {
		d = "erro"
	}
	return d
}

func main() {
	var a, b, c string
	fmt.Scanf("%s\n%s\n%s\n", &a, &b, &c)
	fmt.Println(animal(a, b, c))

}
