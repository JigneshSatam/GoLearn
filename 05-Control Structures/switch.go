package main

import "fmt"

func switchExample() {

	for i := 1; i <= 10; i++ {
		switch i / 2 {
		case 1:
			fmt.Println(i, " is even")
		case 2:
			fmt.Println(i, " is odd")
		default:
			fmt.Println("Nothing")
		}

	}

	add()
	add(1, 2)
	add("i am a ", " string")
	add(1, "dawd", "dawdvwgad", "dawydf", 1.0)
}

func add(anything ...interface{}) {
	switch anything[0].(type) {
	case int:
		switch anything[1].(type) {
		case int:
			fmt.Println("addition is ", anything[0].(int)+anything[1].(int))
		default:
			fmt.Println("Type mismatched error")
		}

	case string:
		switch anything[0].(type) {
		case string:
			fmt.Println("addition is ", anything[0].(string)+anything[1].(string))
		default:
			fmt.Println("Type mismatched error")
		}

	default:
		fmt.Println("Type unknown")
	}

}
