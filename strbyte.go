package main

import (
	"fmt"
)

func main() {
	str1 := "String Byte."
	//b := []byte(str1)

	//for _, x := range b {
	//	fmt.Printf("%v \n", x)
	//}

	//fmt.Println("")

	for x := 0; x < len(str1); x++ {
		if str1[x] == 66 {
			fmt.Print("ok->")
		}
		fmt.Printf("%v\n", str1[x])
	}

}
