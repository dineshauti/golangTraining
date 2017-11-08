package main

import "fmt"

func zero(z *int) {
	*z = 0
	fmt.Println("Address of z: ", z)
}

func main() {

	x := 5
	fmt.Println("Address of x: ", &x)
	fmt.Println("Before zero: ", x)

	zero(&x)
	fmt.Println("After zero: ", x)

}
