// *Funcionou corremente quando foi 'retirado' o 'close(par) e close(impar) da função sendNumbers()'
package main

import "fmt"

func main() {

	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go sendNumbers(par, impar, quit)
	receiveNumbers(par, impar, quit)
}

// envia
func sendNumbers(par, impar chan int, quit chan bool) {

	for x := 0; x <= 20; x++ {
		if x%2 == 0 {
			par <- x
		} else {
			impar <- x
		}
	}
	//close(par)
	//close(impar)
	quit <- true
}

func receiveNumbers(par, impar chan int, quit chan bool) {

	for {
		select {

		case v := <-par:
			fmt.Println("Par:", v)

		case v := <-impar:
			fmt.Println("Impar:", v)

		case v, ok := <-quit:
			if !ok { //se 'ok' for falso
				fmt.Println("err.", v)
			} else {
				fmt.Println("fim.", v)
			}
			return
		}
	}
}
