package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"fmt"
	"os"
)

func main() {
	h := hashset.New()

	s := "dinesh"

	for _,v := range s {

		if !h.Contains(v) {
			h.Add(v)
		} else {
			fmt.Println("The string does not have unique characters")
			os.Exit(3)
		}
	}

	fmt.Println("The string has unique characters")
}
