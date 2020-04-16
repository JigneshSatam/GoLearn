package main

import (
	"fmt"
)

// func main() {
// 	test()

// 	incr2, val2 := initalise()
// 	fmt.Println("line 9", incr2, val2)
// 	// incr2 x => Address 200 ->Value 0
// 	// incr2("incr2")
// 	// incr2 x => Address 200 ->Value 1

// 	// incr("incr1")
// 	// incr1 x => Address 100 ->Value 3
// }

// func test() {
// 	incr, val1 := initalise() // 10000
// 	fmt.Println("line 19", incr, val1)
// 	// incr1 x => Address 100 ->Value 0
// 	incr("incr1")
// 	// incr1 x => Address 100 ->Value 1
// 	incr("incr1")
// 	// incr1 x => Address 100 ->Value 2}
// }

// func initalise() (func(string) int, int) {
// 	x := 0 // 100
// 	y := 1
// 	increment := func(name string) int {
// 		x++ // 100
// 		fmt.Println(name, " is called: ", x)
// 		return x
// 	}
// 	fmt.Println("line 33", increment)
// 	return increment, y
// }

func main() {
	a := initalise()
	a()
	a()
	a()

	b := initalise()
	b()
	a()
	// f2(120)
}

func initalise() func() int {
	x := 0
	f := func() int {
		x++
		fmt.Println("Called: ", x)
		return x
	}
	return f
}

// func f2(x int) int {
// 	fmt.Println("Called: ", x)
// 	return x
// }
