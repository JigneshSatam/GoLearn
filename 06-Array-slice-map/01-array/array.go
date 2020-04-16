package main

import "fmt"

func main() {
	// a = [0,0,0,0,0]
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	// Ruby
	// b << 6

	// interceptor true
	// len b == 5 is full
	// initial new 2 X times the old one
	// copy old array
	// insert new value
	// interceport false

	fmt.Println("dcl:", b)

	c := [4][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	fmt.Println(c)

	// c := [
	// 				[ [1,2,3,4,5], [1,2,3,4,5], [1,2,3,4,5], [1,2,3,4,5] ],
	// 				[ [1,2,3,4,5], [1,2,3,4,5], [1,2,3,4,5], [1,2,3,4,5] ],
	//    ]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[0 1 2] [1 2 3]]
}
