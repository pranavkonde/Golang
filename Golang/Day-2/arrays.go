package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	fmt.Println("Array:", arr1)
	arr2 := [3]int{1, 2, 3}
	fmt.Println("Array:", arr2)
	arr3 := arr2
	arr2[0] = 10
	fmt.Println("Array:", arr2)
	fmt.Println("Array:", arr3)

}
