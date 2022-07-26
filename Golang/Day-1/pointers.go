package main

import "fmt"

func main() {
	a := 12
	var ptr *int = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*ptr)
}
