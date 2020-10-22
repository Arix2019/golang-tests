package main

import (
	"fmt"
)

func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês:", mes)
		fmt.Println("Dia:")
		for dia := 1; dia <= 30; dia++ {
			fmt.Print(" ", dia)
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
