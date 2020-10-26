package main

import "fmt"

func main() {
	sm := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	fmt.Print(sm[0][0])
	fmt.Println(sm[0][3])
	fmt.Print(sm[3][0])
	fmt.Println(sm[3][3])

}
