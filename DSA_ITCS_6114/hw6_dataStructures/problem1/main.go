/*
 *
 * Name: Dinesh Auti
 * Date: 13 Nov 2017
 * Problem 1
 *
 */

package main

import "fmt"

func average(arr []float64) float64 {

	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum/float64(len(arr))
}

func main() {

	var tempArray []float64
	var temp float64
	count := 0

	fmt.Println("Enter six temperature values: ")

	for i:=0; i<6; i++ {
		fmt.Scan(&temp)
		tempArray = append(tempArray, temp)
	}

	avg:= average(tempArray)

	fmt.Println("The average temperature is", avg)

	for index, value := range tempArray {
		if value < avg {
			fmt.Printf("Day %d had temperature %4.2f which was below average \n", index, value)
			count++
		}
	}

	fmt.Printf("The number of days with a temperature below average is %d \n", count)

}


