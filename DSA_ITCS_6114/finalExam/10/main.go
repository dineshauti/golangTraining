package main

import (
	"sort"
	"fmt"
)

var(
	maxTriple int
)
// Complexity: O(nlogn)
func main() {

	input := []int{1, -4, 3, -6, 7, 0}

	sort.Ints(input)

	max1 := input[len(input) - 1] * input[len(input) - 2] * input[len(input) - 3]
	max2 := input[len(input) - 1] * input[0] * input[1]

	if max2 > max1 {
		maxTriple = max2
	} else {
		maxTriple = max1
	}

	fmt.Println(maxTriple)
}
