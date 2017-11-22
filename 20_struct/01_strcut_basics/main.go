package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func (p person) fullName() string {
	return p.firstName +" "+ p.lastName
}

func main() {
	p1 := person{"James", "Bond",20}

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p1.fullName())
}
