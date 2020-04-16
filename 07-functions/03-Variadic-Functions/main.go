package main

import (
	"fmt"
)

func add(x ...int) int {
	total := 0
	// fmt.Println(args[0])
	for _, v := range x {
		total += v
	}
	return total
}

// func returnFunc() func(string, int) int {
// 	f := func(str string, in int) int {
// 		fmt.Println("In the inner function")
// 		return 0
// 	}
// 	return f
// }

func main() {
	// fmt.Println("With 0 args sum = ", add())
	// fmt.Println("With 1 args sum = ", add(1))
	// fmt.Println("With 2 args sum = ", add(1, 2))
	// fmt.Println("With 3 args sum = ", add(1, 2, 3))
	x := returnFunc()
	a := x(12)
	fmt.Println("a: ", a)

}

func returnFunc() func(int) string {
	x := func(i int) string {
		fmt.Println(i)
		return "Hi"
	}
	return x
}
