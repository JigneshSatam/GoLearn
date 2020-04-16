package main

import (
	"fmt"
)

func main() {

	// forFunc()

	// ifelse()

	switchExample()
}

func forFunc() {
	for {
		fmt.Println("Never Ends")
	}

	for i := 0; i < 11; i++ {
		fmt.Println(" Hello, playground ", i)

	}

	arr := []string{"Hello", "PPL"}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println("Index is ", i, " value is ", v)
	}

	for _, v := range arr {
		fmt.Println("Value is ", v)
	}
}
