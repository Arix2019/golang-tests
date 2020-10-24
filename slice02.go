package main

import (
	"fmt"
)

func main() {
	nomes := []string{"Ana", "Michelle", "Patricia", "Helen"}

	nomes = append(nomes, "Rosana")

	for i, v := range nomes {
		//fmt.Println(v, "encontra-se no índice", i)
		fmt.Printf("O nome %v encontra-se no indice %v\n", v, i)
	}
}
