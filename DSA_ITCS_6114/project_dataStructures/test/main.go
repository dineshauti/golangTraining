package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var(
	inputArray []float64
)


func main() {

	//reads line all in one go
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	arr := strings.Split(input, " ")

	for _,x := range arr {

		y,_ := strconv.ParseFloat(strings.TrimSpace(x),64)
		inputArray = append(inputArray, y)
	}

	fmt.Println(inputArray)


}
