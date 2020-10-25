// deletando elementos de um slice
package main

import (
	"fmt"
)

func main() {

	frutas := []string{"Banana", "Maça",
		"Laranja", "Uva"}

	/*	e := 2 // atribui variavel 'e' a posição do elemento que será removido.

		copy(frutas[e:], frutas[e+1:])  // Shift a[i+1:] left one index.
		frutas[len(frutas)-1] = ""     // Erase last element (write zero value).
		frutas = frutas[:len(frutas)-1] // Truncate slice.
	*/
	copy(frutas[2:], frutas[2+1:])
	frutas[len(frutas)-1] = ""
	fmt.Println(frutas)

}
