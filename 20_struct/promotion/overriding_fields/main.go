package main

import "fmt"

type person struct {
	First string
	Last string
	Age int
}

type doubleZero struct {
	person
	First string
	toKill bool
}


func main() {
	p1 := doubleZero{
		person{
			First:"Dinesh",
			Last:"Auti",
			Age:27,
		},
		"Maianne",
		false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.toKill)
	fmt.Println(p1.person.First, p1.Last, p1.Age, p1.toKill)
}
