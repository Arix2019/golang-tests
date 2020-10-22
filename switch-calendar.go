package main

import (
	"fmt"
)

func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Printf("Mês:%v\n", mes)
		switch {
		case mes == 12:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 11:
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 10:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 9:
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 8:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		case mes%2 == 0:
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}
		case mes%2 != 0:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		}
		fmt.Println(" ")
	}
}
