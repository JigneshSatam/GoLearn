package main

import (
	"fmt"
)

// func toStr(x int, y int) (string, int, []int) {
// 	str := strconv.Itoa(x) + strconv.Itoa(y)
// 	val := x + y
// 	return str, val, []int{x, y}
// }

func main() {
	// str1, int1, arr := toStr(5, 10)
	// fmt.Println("Concatenated Strings: ", str1, int1, arr)
	a, b := swap(5, 10)
	fmt.Println(a, b)
}

func swap(x int, y int) (int, float32) {
	return y, (float32(x) / 2.0)
}
