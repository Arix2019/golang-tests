package main

import (
	"fmt"
)

func main() {
	x := 3

	if x > 6 {
		fmt.Println(x, "-> maior que 6.")
	} else if x < 6 {
		fmt.Println(x, "-> menor que 6.")
	} else {
		fmt.Println("Igual.")
	}

}
