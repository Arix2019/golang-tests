// exibe conteúdo em formato json
package main

import (
	"encoding/json"
	"fmt"
)

type Clube struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Fundacao int    `json:"fundacao"`
	Pais     string `json:"pais"`
}

var clubes = []Clube{
	Clube{
		Id:       1,
		Nome:     "Bayer de Munique",
		Fundacao: 1900,
		Pais:     "Alemanha",
	},
	Clube{
		Id:       2,
		Nome:     "Paris Saint German",
		Fundacao: 1974,
		Pais:     "França",
	},
}

func convJson() {
	js, err := json.Marshal(clubes) // converte para o formato json
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}

func main() {
	convJson()
}
