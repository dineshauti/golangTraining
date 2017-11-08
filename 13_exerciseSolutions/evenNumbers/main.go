package main

import "fmt"

func main() {

	//Printing even numbers
	for i := 0; i<100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
