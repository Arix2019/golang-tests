/*
*	utilizando canais
 */
package main

import (
	"fmt"
)

var channel = make(chan int)

func main() {

	go send(channel)
	receive(channel)

}

func send(s chan<- int) { //valor '51' sendo atribuido ao canal (channel).
	s <- 51
}

func receive(r <-chan int) {
	fmt.Println("O valor recebido pelo canal foi:", <-r)
}
