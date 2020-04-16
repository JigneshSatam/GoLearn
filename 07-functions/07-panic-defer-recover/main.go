package main

import "fmt"

func main() {
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	defer defered()
	panic("Error")
}

// func main() {
// 	defer defered()
// 	panic("PANIC")
// }

func defered() {
	str := recover()
	fmt.Println(str)
}
