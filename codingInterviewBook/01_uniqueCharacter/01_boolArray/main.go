package _1_boolArray

import (
	"fmt"
	"os"
)


func main() {

	s := "dinesh"
	boolSlice := make([]bool, 128)

	for _,v := range s {

		if boolSlice[v] == false {
			boolSlice[v] = true
		} else {
			fmt.Println("The string does not have unique characters")
			os.Exit(3)
		}
	}

	fmt.Println("The string has unique characters")
}
