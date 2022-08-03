package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice1)
	slice1 = slice1[2:5]
	// slice2 := append(slice1, 9, 1, 1, 2, 3, 4, 5, 6)
	// fmt.Println(slice2)
	// slice1[2] = 90
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// slice2 := append(slice1, 9)
	// fmt.Println(slice2)
	// slice1[2] = 90
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	slice2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)
}
