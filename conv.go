package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//int para string
	num := 12
	fmt.Printf("O valor (%v) é do tipo: (%T).\n", num, num)
	x := strconv.Itoa(num)
	fmt.Printf("Agora (%v) passou a ser do tipo: (%T).\n", x, x)
	//string para int
	y, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mas (%v) voltou a ser do tipo: (%T).\n", y, y)
}
