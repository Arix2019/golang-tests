package main

import "fmt"

func main() {

	x := 0
	for x <= 10 {
		defer fmt.Println(x)
		x++
	}
}
