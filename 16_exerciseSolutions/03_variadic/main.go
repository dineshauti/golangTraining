package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for i,v := range numbers {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(max(-10, -2))
}
