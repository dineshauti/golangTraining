package main

import "fmt"

func countBinarySubstrings(s string) int {
	prevRunLength := 0
	currRunLength := 1
	res := 0

	for i:=1; i<len(s); i++ {

		if  s[i] == s[i-1] {
			currRunLength++
		} else {
			prevRunLength = currRunLength
			currRunLength = 1
		}

		if prevRunLength >= currRunLength {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countBinarySubstrings("110001111000000"))
}
