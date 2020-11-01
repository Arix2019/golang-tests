package main

import (
	"fmt"
)

// strucs
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	fone      int
}

type atividade struct {
	pessoa
	profissao string
	local     string
	salario   float64
}

type filhos struct {
	pessoa
	qtd int
}

type filhosPessoaAtividade struct {
	pessoa
	atividade
	qtd int
}

type filhosMenoresIdade struct {
	pessoa
	atividade
	filhos
	qtdFilhosMenores int
}

// strucs

var cadastro01 = atividade{
	pessoa: pessoa{
		nome:      "Maria",
		sobrenome: "Joaquina",
		idade:     23,
		fone:      99564326,
	},
	profissao: "vendedora",
	local:     "loja",
	salario:   3456.50,
}

var cadastro02 = pessoa{
	nome:      "Rosa",
	sobrenome: "Miller",
	idade:     29,
	fone:      99654876,
}

var cadastro03 = filhos{
	pessoa: pessoa{
		nome:      "Lucia",
		sobrenome: "Lopez",
		idade:     32,
		fone:      98875437,
	},
	qtd: 3,
}

var cadastro04 = filhosMenoresIdade{
	pessoa: pessoa{
		nome:      "Lucia",
		sobrenome: "Lopez",
		idade:     32,
		fone:      98875437,
	},
	atividade: atividade{
		profissao: "Arquiteta",
		local:     "Linux Arq.",
		salario:   8567.44,
	},
	filhos: filhos{
		qtd: 3,
	},
	qtdFilhosMenores: 1,
}

var cadastro05 = pessoa{
	nome:      "Erika",
	sobrenome: "Figueiredo",
	idade:     40,
	fone:      9956255414,
}

func main() {

	fmt.Println(cadastro01)
	fmt.Println(cadastro02)
	fmt.Println(cadastro03)
	fmt.Printf("Nome: %v\nSobrenome: %v\nIdade: %v\nFone: %v\n",
		cadastro05.nome, cadastro05.sobrenome, cadastro05.idade, cadastro05.fone)
	fmt.Println("")
	fmt.Println("Nome:", cadastro04.nome,
		"\nSobrenome:", cadastro04.sobrenome,
		"\nProfissão:", cadastro04.profissao,
		"\nSalário:", cadastro04.salario,
		"\nFilho(s):", cadastro04.qtd,
		"\nFilho(s) menor(es) de idade:", cadastro04.qtdFilhosMenores)
}
