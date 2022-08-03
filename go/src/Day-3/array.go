package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	fmt.Println("Arr1", arr1)
	arr2 := [3]int{1, 2, 3}
	fmt.Println("Arr1", arr2)
	fmt.Printf("Type of arr is  %T", arr1) // print size and data type too
	fmt.Println(arr1 == arr2)
	arr3 := [3][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)
}
