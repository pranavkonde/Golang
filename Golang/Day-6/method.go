package main

import "fmt"

type Receiver int

func (r Receiver) someFunc() {
	r = 10
	fmt.Println("R:", r)
}
func main() {
	var r1 Receiver
	r1.someFunc()
}
