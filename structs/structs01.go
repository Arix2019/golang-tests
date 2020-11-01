// Structs permite armazenar dados de tipos
// 	diferentes (int,string,float...)
package main

import "fmt"

type info struct {
	nome   string
	idade  int
	filhos bool
}

func main() {

	// 1º método
	info1 := info{
		nome:   "Martina",
		idade:  24,
		filhos: false,
	}

	// 2º método
	info2 := info{"Roberta", 30, true}

	fmt.Println("Nome:", info1.nome)
	fmt.Println(info2)

}
