/* 								gera um Hash para senha
   foi necessário a instalação do pacote: (go get -u golang.org/x/crypto/bcrypt)
   Equivale ao custo computacional em Gerar os hashes: '4' o menor valor / '31' o maior e 10 default.
const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPa
)
*/
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwd := "Orinocoflow"
	passfail := "OrinocoFrow"

	sb, err := bcrypt.GenerateFromPassword([]byte(passwd), 10) // gera o hash
	if err != nil {
		fmt.Println(err)
	}

	passok := bcrypt.CompareHashAndPassword(sb, []byte(passwd))      //compara o 'slice of byte' com a senha correta.
	wrongpass := bcrypt.CompareHashAndPassword(sb, []byte(passfail)) //compara com a senha errada.

	fmt.Println("Hash:", string(sb))
	fmt.Println("Bingo!", passok)
	fmt.Println("Fail:", wrongpass)

}
