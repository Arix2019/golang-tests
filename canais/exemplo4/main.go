package main

import "fmt"

func main() {
	canal := make(chan string)

	go enviaTexto("Olar", canal)
	recebeTexto(canal)

}

func enviaTexto(text string, e chan<- string) { //'n' é a variavel - 'e chan' á a variavel que enviará o resultado da func
	e <- text
}

func recebeTexto(r <-chan string) {

	fmt.Printf("O conteúdo do canal recebido foi: '%v' e é do tipo: '%T'\n", <-r, r)
}
