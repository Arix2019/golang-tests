package main

import (
	"fmt"
)

func main() {

	c := make(chan int)    //canal bidirecional
	cs := make(chan<- int) //canal envia
	cr := make(<-chan int) //canal recebe

	fmt.Printf("Canal bidirecinal: %T\n", c)
	fmt.Printf("Send channel: %T\n", cs)
	fmt.Printf("Receive channel: %T\n", cr)

	fmt.Printf("\n---Geral para especifico---\n")
	fmt.Printf("c  %T\n", (<-chan int)(c))
	fmt.Printf("c  %T\n", (chan<- int)(c))

}
