package main

import "fmt"

func main() {
	gainDBConnection()
	defer closeDBConn()
}

func gainDBConnection() {
	// panic("Someting went wrong")
	fmt.Println("Writing into DB")
}
func returnDBConnection() {
	fmt.Println("Closing DB")
}

func closeDBConn() {
	returnDBConnection()
	func() {
		fmt.Println("returnDBConnection2")
	}()
}
