package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Marmelade"
	str2 := "marmelade"
	str3 := "Marmelade"

	fmt.Println(strings.Compare(str, str2))
	fmt.Println(strings.Compare(str, str3))

}
