package main

import (
	"fmt"
)

func main() {
	n := 20

	fmt.Printf("Base binária do valor (%v): %b\n", n, n)
	fmt.Printf("Base octal do valor (%v): %o\n", n, n)
	fmt.Printf("Base hexadecimal do valor (%v): %x\n", n, n)
}
