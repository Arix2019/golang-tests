package main

import "fmt"

func main() {
	sm1 := [][][][]int{

		[][][]int{
			[][]int{
				[]int{1, 2, 3},
			},
			[][]int{
				[]int{4, 5, 6},
			},
		},
		[][][]int{
			[][]int{
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{10, 11, 12},
			},
		},
	}

	fmt.Println(sm1)
	//fmt.Println(len(sm1))
}
