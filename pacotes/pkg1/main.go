/*
* pacotes pata organizar o código
* pkg1 - main.go
* Quando houver uma ou mais funções que estão em outro arquivo o comando
* para compilar a função 'main' será: $ go run *.go
 */
package main

import (
	"fmt"
	"github.com/golang-tests/pacotes/pkg2"
)

func main() {

	fmt.Println("Função 'main' principal (pkg1).")
	SecondFunc()
	pkg2.Pkg2() // importa a função 'Pkg2()' de dentro do pacote 'pkg2'.
}
