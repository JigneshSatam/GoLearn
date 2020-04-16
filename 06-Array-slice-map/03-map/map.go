package main

import "fmt"

func main() {
	// m1 := make(map[int]string)
	// for i := 0; i < 3; i++ {
	// 	m1[i] = "hah"
	// }
	// fmt.Println("map:", m1)
	// m1 := map[int]string{0: "ws"}
	// m1[3] = "Yo"

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 0

	fmt.Println("map:", m)

	v1 := m["k3"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	v2, ok := m["k2"]
	fmt.Println("value:", v2)
	fmt.Println("prs:", ok)

	n := map[string]int{"foo": 1, "bar": 2, "fao": 3}
	fmt.Println("map:", n)
}
