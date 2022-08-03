package main

import "fmt"

func main() {
	var a int
	var b int = 1
	fmt.Scanln(&a)
	for i := 2; i <= a; i++ {
		if a%i == 0 {
			b++
		}
	}
	if b > 2 {
		fmt.Println("Not Prime")
	} else {
		fmt.Println("Prime number")
	}
}
