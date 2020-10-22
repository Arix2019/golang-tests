package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 12; x++ {
		if x%2 == 0 {
			fmt.Println("Par:", x)
		} else {
			fmt.Println("Impar:", x)
		}
	}
}
