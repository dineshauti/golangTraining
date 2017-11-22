package main

import "fmt"

func main() {

	pr1,pr2,pr3,pr4,pr5,pr6,pr7,pr8 := 1.0,1.0,1.0,1.0,1.0,1.0,1.0,1.0

	for i:=0; i<1000; i++ {

		pr1 = 0.15+0.85*(pr3/2)
		pr2 = 0.15+0.85*(pr5/2 + pr7/2)
		pr3 = 0.15
		pr4 = 0.15+0.85*(pr5/2)
		pr5 = 0.15+0.85*(pr3/2)
		pr6 = 0.15+0.85*(pr2 + pr4)
		pr7 = 0.15+0.85*(pr1)
		pr8 = 0.15+0.85*(pr6 + pr7/2)
	}

	fmt.Println(pr1)
	fmt.Println(pr2)
	fmt.Println(pr3)
	fmt.Println(pr4)
	fmt.Println(pr5)
	fmt.Println(pr6)
	fmt.Println(pr7)
	fmt.Println(pr8)

}
