// TIPOS DE VARIÁVEIS
package main

import (
	"fmt"
)

func main() {
	i := 123
	f := 45.1
	c := 0.154 + 0.5i
	s := "Sou uma string."
	b := true

	fmt.Printf("A variável (%v) é do tipo: %T\n", i, i)
	fmt.Printf("A variável (%v) é do tipo: %T\n", f, f)
	fmt.Printf("A variável (%v) é do tipo: %T\n", c, c)
	fmt.Printf("A variável (%v) é do tipo: %T\n", s, s)
	fmt.Printf("A variável (%v) é do tipo: %T\n", b, b)
}
