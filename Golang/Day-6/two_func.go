package main

import "fmt"

type Employee struct {
	ID   int
	name string
}

func (e Employee) renameEmployee() {
	e.ID = 10
	e.name = "Pranav"
}
func renameEmployeeFunc(e Employee) {
	e.ID = 100
	e.name = "Piyush"
}
func main() {
	e1 := Employee{
		ID:   0,
		name: "Initial name",
	}
	e1.renameEmployee()
	fmt.Println("after method:", e1)
	renameEmployeeFunc(e1)
	fmt.Println("After method:", e1)
}
