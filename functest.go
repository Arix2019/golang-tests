package main

import (
	"fmt"
)

func calc(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("O valor da função é:", calc(45, 21))
}
