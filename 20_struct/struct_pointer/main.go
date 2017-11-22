package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p1 := person{"Dinesh", 27}
	p2 := &person{"Marianne", 25}

	fmt.Println(p1.name)
	fmt.Println(p2.age)
	fmt.Printf("%T\n", p1)
}
