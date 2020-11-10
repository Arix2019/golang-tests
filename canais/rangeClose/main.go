/*
*	Comunicação entre funções (canais)
 */
package main

import "fmt"

func main() {
	c := make(chan int)

	go loop(10, c)
	show(c)
}

func loop(cont int, s chan<- int) {
	for x := 0; x <= cont; x++ {
		s <- x
	}
	close(s)
}

func show(r <-chan int) {
	for x := range r {
		fmt.Println("Valor recebido do canal:", x)
	}
}
