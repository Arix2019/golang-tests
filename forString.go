package main

import (
	"fmt"
)

func main() {

	x := 0

	for {

		if x <= 10 {
			fmt.Printf("%v->%b\n", x, x, x)

		} else {
			break
		}

		x++
	}
}
