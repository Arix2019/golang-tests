package main

import (
	"fmt"
)

func main() {

	fmt.Println("O valor da função é:", calc(45, 21))
}

func calc(x, y int) int {
	return x + y
}
