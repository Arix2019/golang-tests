// o primeiro 'defer' sempre será executado por último
package main

import "fmt"

func main() {

	defer fmt.Println("Será exibido em 1º.")
	fmt.Println("Será exibido em 2º.")
	fmt.Println("Será exibido em 3º.")
}
