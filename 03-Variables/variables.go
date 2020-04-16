package main

import (
	"fmt"
)

type user struct {
	name string
}

func main() {
	var r = "a"

	r = "b"

	var defaultUsr user
	fmt.Println(defaultUsr)

	var str string
	str = "adwd"
	fmt.Println(str)

	fmt.Println(r)

	a := "1"

	fmt.Println(a)

	x, y := 1, "2.1"
	// x, y = y, x
	fmt.Println(x, y)

}
