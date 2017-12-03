package main

import "fmt"

type square struct {
	Side float64
}

type circle struct {
	Radius float64
}

// Circle implements a shape interface
func (c circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Square implements a shape interface
func (s square) area() float64 {
	return s.Side * s.Side
}


// Interface is an type - IMPORTANT
// Interface can be used to exhibit polymorphism
type shape interface {
	area() float64 // This defines how the interface needs to be implemented and anything that implements that method is
	               // said to have implemented the interface
}

func info(z shape) { // info takes in a shape interface
	fmt.Println(z)
	fmt.Println(z.area())
}


func main() {

	s := square{10}
	info(s)

	c := circle{10}
	info(c)
}
