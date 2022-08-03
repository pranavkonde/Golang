package main

import "fmt"

func main() {

	func() {
		fmt.Println("Inside anonymous function")
	}()
}

// s:=func() {
// 	fmt.Println("Inside anonymous function")
// }
// s()
// }
