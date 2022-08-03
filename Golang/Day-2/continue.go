package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd no
			continue
		}
		sum += i
	}
	fmt.Println(sum)

}
