// channel / convergencia
package main

import (
	"fmt"
	"sync"
)

func main() {

	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go enviar(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valores Recebidos:", v)
	}
}

func enviar(p, i chan int) {

	for x := 0; x <= 20; x++ {
		if x%2 == 0 {
			p <- x
		} else {
			i <- x
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}
