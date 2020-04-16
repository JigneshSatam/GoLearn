package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	for i := 1.0; i <= 10000000; i++ {
		fact(170.0)
	}
	fmt.Println()
	fmt.Println("============================")
	fmt.Println("======", time.Since(t1), "=======")
	fmt.Println("============================")

}

func fact(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}
