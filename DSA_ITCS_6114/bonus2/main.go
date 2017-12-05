/* Title: Compute the power of a matrix
 * Name: Dinesh Auti
 * Date: 4 Dec, 2017
*/

package main

import (
	"gonum.org/v1/gonum/mat"
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var(
	dim int
	power int
	inputArray []float64
)


func main() {

	fmt.Print("Enter the dimensions of the matrix: ")
	fmt.Scanf("%d",&dim)

	fmt.Print("Enter the elements of the matrix: ")

	//reads line all in one go
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputSplit := strings.Split(input, " ")

	for _,x := range inputSplit {
		y,_ := strconv.ParseFloat(strings.TrimSpace(x),64)
		inputArray = append(inputArray, y)
	}

	fmt.Print("Enter the power that you would like to compute: ")
	fmt.Scanf("%d",&power)

	m1 := mat.NewDense(dim,dim, inputArray)
	m1.Pow(m1,power)
	fmt.Printf("m :\n%f\n\n", mat.Formatted(m1, mat.Prefix(""), mat.Excerpt(2)))
}
