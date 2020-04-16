package main

import "fmt"

func ifelse() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else if i%2 == 1 {
			fmt.Println(i, " is odd")
		} else {
			fmt.Println("Nothing")
		}
	}
}
