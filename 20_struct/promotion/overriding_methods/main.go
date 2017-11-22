package main

import "fmt"

type person struct {
	Name string
	Age int
}

type doubleZero struct {
	person
	toKill bool
}

func (p person) Greeting() {
	fmt.Println("I am fucking awesome!")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Good to see you!")
}

func main() {

	p1 := person{
		"Dinesh",
		27,
	}

	p2 := doubleZero{
		person{
			Name: "Marianne",
			Age: 25,
		},
		false,
	}

	p2.Greeting()
	fmt.Println(p1.Name)

}
