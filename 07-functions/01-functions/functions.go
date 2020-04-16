package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c string) string {
	return a + b + c
	// return 0
	// return "return string"
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res1 := plusPlus("Hell", "o", "World")
	fmt.Println("1+2+3 =", res1)
}
