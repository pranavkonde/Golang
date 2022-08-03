package main

import "fmt"

func main() {
	var s, s1 string
	fmt.Scanln(&s)
	s1 = (longestPaldrome(s))
	fmt.Println(s1)
}
func longestPaldrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		s1 := s[i:]
		for j := 0; len(s1) > len(longest) && j < len(s1); j++ {
			s2 := s1[:j+1]
			if len(s2) > len(longest) && getReverse(s2) {
				longest = s2
			}
		}
	}
	return longest
}
func getReverse(s2 string) bool {
	j := len(s2) - 1
	for i := 0; i < len(s2); i++ {
		if s2[i] != s2[j] {
			return false
		}
		j--
	}
	return true
}
