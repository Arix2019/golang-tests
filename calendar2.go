package main

import (
	"fmt"
)

// mes 8 -> 31 dias
func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês:", mes)
		fmt.Println("Dia:")
		if mes == 8 {
			for dia2 := 1; dia2 <= 31; dia2++ {
				fmt.Print(" ", dia2)
			}
		} else if mes%2 == 0 {
			for dia2 := 1; dia2 <= 30; dia2++ {
				fmt.Print(" ", dia2)
			}

		} else {
			for dia2 := 1; dia2 <= 31; dia2++ {
				fmt.Print(" ", dia2)
			}
		}
		fmt.Println("\n")
	}

}

/*
Janeiro-->  31 dias
Fevereiro-->  tem 28 dias e 29 em anos bissextos
Março-->  31 dias
Abril--> 30 dias
Maio-->  31 dias
Junho-->  30 dias
Julho-->  31 dias
Agosto-->  31 dias
Setembro-->  30 dias
Outubro-->  31 dias
Novembro-->  30 dias
Dezembro-->  31 dias
*/
