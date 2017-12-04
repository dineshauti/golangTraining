//Accepted
package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {

	arr := make([]string, n)

	for i:=0; i<n; i++ {
		if (i+1)%15 == 0{
			arr[i] = "FizzBuzz"
		} else if (i+1)%3 == 0 {
			arr[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			arr[i] = "Buzz"
		} else {
			arr[i] = strconv.Itoa(i+1)
		}
	}
	return arr
}

func main() {
	fmt.Println(fizzBuzz(5))
}
