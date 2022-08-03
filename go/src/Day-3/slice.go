package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr1[2:6])
	slice2 := arr1[3:8]
	slice3 := arr1[0:3]
	fmt.Println(slice2)
	fmt.Println(slice3)
	arr1[5] = 50
	fmt.Println("slice2: capacity: length: ", slice2, cap(slice2), len(slice2))

}
