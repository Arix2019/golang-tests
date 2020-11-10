package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mainChannel := converge(trabalho("-First"), trabalho("Second"))

	for x := 0; x < 20; x++ {
		fmt.Println(<-mainChannel)
	}
}

//criar uma função 'trabalho' que tenha uma canal e que retorne dados para o canal:
func trabalho(s string) chan string {
	channel := make(chan string)
	go func(s string, ch chan string) {
		for x := 0; ; x++ {
			ch <- fmt.Sprintf("A função %v diz: %v", s, x)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) //vai gerar um tempo aleatório entre 0 e 1 segundo.
		}
	}(s, channel)

	return channel
}

func converge(x, y chan string) chan string {
	novoCanal := make(chan string)
	go func() {
		for {
			//canal_novo' vai receber td que sair do canal 'x'
			novoCanal <- <-x
		}
	}()
	go func() {
		for {
			//'canal_novo' vai receber td que sair do canal 'y'
			novoCanal <- <-y
		}
	}()
	return novoCanal
}
