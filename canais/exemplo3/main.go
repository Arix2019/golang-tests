/*
	'Canais' Permitem transmitir dados
	entre goroutines.
*/
package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Olar."
	}()

	fmt.Println(<-channel)
}
