package main

import "fmt"

func main() {

	canal := make(chan int)

	go func() {
		canal <- 51
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok)
}
