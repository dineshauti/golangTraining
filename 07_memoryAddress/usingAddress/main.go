package main

import "fmt"

const metersToYard float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enters meters swam: ")
	fmt.Scan(&meters)

	yards := meters * metersToYard

	fmt.Println(meters, "meters is ", yards, "in yards")
}

