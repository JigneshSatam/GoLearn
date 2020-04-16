package main

import "fmt"

func main() {
	// a := []int{1, 2}

	// fmt.Println("initial len of a: ", len(a), a)
	// x := append(a, 2)
	// a[0] = 20
	// fmt.Println("Updated len of a: ", len(a), a, x)

	s := make([]string, 5)
	fmt.Println("emp:", s)
	// ["" "" "" "" ""] 101 102.0 1000

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	s[3] = "h"
	s[4] = "z"
	//OLD s =>Address = 100 -> ["a", "b", "c", "h", "z"]
	s = append(s, "d")
	//NEW s =>Address = 200 -> ["a", "b", "c", "h", "z", "d"]
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	x := s // Pointer/Reference assignment
	// s and x => Address = 200 -> ["a", "b", "c", "h", "z", "d"]
	x[0] = "x"
	fmt.Println("modified:", s)

	c := make([]string, len(s))
	// s and x => Address = 200 -> ["a", "b", "c", "h", "z", "d"]
	copy(c, s)
	// s and x => Address = 200 -> ["a", "b", "c", "h", "z", "d"]
	// c => Address = 300 -> ["a", "b", "c", "h", "z", "d"]
	c[0] = "c"
	fmt.Println("old:", s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("s:", s)
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("s:", s)
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("s:", s)
	fmt.Println("sl3:", l)

	arr := [2]string{"Hello", "World"}
	l = arr[:] // Pointer/Reference assignment
	l[1] = "New"
	fmt.Println("arr:", arr)
	fmt.Println("sl4:", l)

	l[0] = "l"
	fmt.Println("s:", s)
	fmt.Println("modified l :", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
