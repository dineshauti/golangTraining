package main

import (
	"fmt"
)


func main() {

	a := make([]map[int]int, 3)
	for i := range a {
		a[i] = make(map[int]int, 5)
	}

	a[2][20] = 3434

	fmt.Println(a)


}
