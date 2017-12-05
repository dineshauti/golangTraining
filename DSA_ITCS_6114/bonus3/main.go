/* Title: Compute the power of a matrix
 * Name: Dinesh Auti
 * Date: 4 Dec, 2017
*/

package main

import (
	"math"
	"fmt"
)

var(
	factorsList []int
	num int
)

func factors(num int) {

	// Print the number of 2s that divide num
	for num%2 == 0 {
		factorsList = append(factorsList, 2)
		num = num /2
	}

	for i:=3; i<=int(math.Sqrt(float64(num))); i=i+2 {
		for num%i == 0 {
			factorsList = append(factorsList, i)
			num = num/i
		}
	}

	// If prime include the number
	if num >2 {
		factorsList = append(factorsList, num)
	}
}


func main() {

	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d",&num)

	factors(num)

	factorsMap := make(map[int]int)

	for _,x := range factorsList {
		if val,ok := factorsMap[x]; ok {
			factorsMap[x] = val+1
		} else {
			factorsMap[x] = 1
		}
	}

	n := 1

	for i,y := range factorsMap {
		if y%2 != 0 {
			n = n*i
		}
	}

	fmt.Println("n: ",n)
}
