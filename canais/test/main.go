package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	end := make(chan bool)

	go sendText("OIE!!!", ch)
	go receiveText(ch, end)
	<-end
}

func sendText(str string, e chan string) {
	e <- str
}

func receiveText(r chan string, end chan bool) {
	fmt.Println("Recebido:", <-r)
	end <- true
}
