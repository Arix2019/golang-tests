/*
	agrupamento de dados - Array - Slice
*/
package main

import (
	"fmt"
)

func main() {
	mes := []string{"Janeiro", "Fevereiro",
		"Março", "Abril", "Maio",
		"Junho", "Julho", "Agosto",
		"Stembro", "Outubro", "Novembro", "Dezembro"}

	for i, v := range mes {
		fmt.Println(v)
		if i%2 == 0 {
			for dia := 0; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}

		} else {
			for dia := 0; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}

		}
		fmt.Println(" ")
	}
}
