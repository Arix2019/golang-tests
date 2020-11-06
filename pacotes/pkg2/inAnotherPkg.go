//pkg2 - inAnotherPkg.go
package pkg2 //nome do pacote que será expotado

import "fmt"

func Pkg2() {
	fmt.Println("Essa função foi importada de outro pacote (pkg2)!")
}
