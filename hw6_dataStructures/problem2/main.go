/*
 *
 * Name: Dinesh Auti
 * Date: 13 Nov 2017
 * Problem 2
 *
 */

package main

import "fmt"

func main() {
	var numElements int
	var num float64
	var arr []float64

	fmt.Println("How many numbers will you enter?")
	fmt.Scan(&numElements)

	fmt.Printf("Enter %d numbers one per line \n", numElements)

	for i:=0; i<numElements; i++ {
		fmt.Scanln(&num)
		arr = append(arr, num)
	}

	fmt.Printf("You have filled the array with %d numbers. \n", numElements)

	sum := 0.0
	product := 1.0

	for _, value := range arr {
		sum += value
		product *= value
	}

	fmt.Println("The sum of the numbers =", sum)
	fmt.Println("The product of the numbers =", product)

	fmt.Println("Here are the numbers and their percent contribution to the sum.")

	for _, value := range arr {
		fmt.Printf("%f is %f%% of the sum. \n", value, (value/sum)*100)
	}

}
