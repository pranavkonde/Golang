package main

import "fmt"

type Receive struct {
	ID   int
	name string
}

func (r Receive) someFunc() {
	r.ID = 20
	r.name = "Pranav"
	fmt.Println("Inside Function: ", r)
}
func main() {
	r1 := Receive{
		ID:   200,
		name: "Piyush",
	}
	r1.someFunc()
	fmt.Println("Inside main sFunction: ", r1)
}
