package main

import "fmt"

// função variadica (...)
func soma(x ...int) int {

	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func main() {

	sl := []int{5, 10, 54, 31, 1, 20}

	fmt.Println(soma(sl...))
}
