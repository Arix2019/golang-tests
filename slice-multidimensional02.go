package main

import "fmt"

func main() {
	sm1 := [][][][]int{

		// 0
		[][][]int{
			[][]int{
				[]int{1, 2, 3},
			},
			[][]int{
				[]int{4, 5, 6},
			},
		},
		// 1
		[][][]int{
			[][]int{
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{10, 11, 12},
			},
		},
	}

	fmt.Print(sm1[0][0][0][0], "-") //result: 1
	fmt.Print(sm1[0][0][0][1], "-") //result: 2
	fmt.Println(sm1[0][0][0][2])    //result: 3

	fmt.Println(sm1[0][1][0][0]) //result: 4
	fmt.Println(sm1[1][0][0][0]) //result: 7
	fmt.Println(sm1[1][1][0][0]) //result: 10

	/*	fmt.Println(sm1[0][0])
		fmt.Println(sm1[0][1])
		fmt.Println(sm1[1][0])
		fmt.Println(sm1[1][1])  */
}
