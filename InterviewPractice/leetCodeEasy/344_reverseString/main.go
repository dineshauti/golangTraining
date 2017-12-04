package main

import "fmt"

func reverseString(s string) string {

	// Get Unicode code points.
	rune := []rune(s)
	n := len(rune)

	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	// Convert back to UTF-8.
	output := string(rune)

	return output
}

func main() {
	fmt.Println(reverseString("Hello"))
}
