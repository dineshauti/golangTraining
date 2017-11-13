package main

import (
	"fmt"
	"math"
)

func Primes(low, high int) {
	total := 0
	f := make([]bool, high+1)
	for i := 2; i <= int(math.Sqrt(float64(high))); i++ {
		if f[i] == false {
			for j := i * i; j <= high; j += i {
				f[j] = true
			}
		}
	}
	for i := low; i <= high; i++ {
		if f[i] == false {
			if i != 0 && i != 1 {
				fmt.Printf("%v ", i)
				total++
			}
		}
	}
	fmt.Println("")
	fmt.Printf("There are %d primes. \n", total)
}

func main() {

	var num int
	var arr []int

	fmt.Println("Enter two positve integers (in increasing order):")

	for i:=0; i<2; i++ {
		fmt.Scanln(&num)
		arr = append(arr, num)
	}

	fmt.Printf("Prime numbers between %d and %d \n", arr[0], arr[1])

	Primes(arr[0], arr[1])

}
