/*
* - Criar dois canais
* - Uma func manda X números ao primeiro canal
* - Outra func faz um range deste canal, e para cada item lança uma go func que coloca o retorno de trabalho() no canal dois
* - Trabalho é um timer aleatorio para simular workload
* Por fim, range canal dois demonstra os valores.
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go sendNumbers(22, c1)
	go receiveNumbers(c1, c2)

	for v := range c2 {
		fmt.Println("Channel two range:", v)
	}

}

// envia um range de números para o primeiro canal ('c1')
func sendNumbers(n int, sendC1 chan int) {
	for x := 0; x <= n; x++ {
		sendC1 <- x
	}
	close(sendC1)
}

//(2º canal) recebe o valor passado pela func sendNumbers() - (1ºcanal)
func receiveNumbers(receiveC1, receiveC2 chan int) {
	var wg sync.WaitGroup
	for v := range receiveC1 {
		wg.Add(1)
		go func(num int) {
			receiveC2 <- trabalho(num)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(receiveC2)
}

//envia uma sequencia aleatoria para func receiveNumbers()
func trabalho(t int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
	// retorna o range de numeros da função sendNumbers()
	return t
}
