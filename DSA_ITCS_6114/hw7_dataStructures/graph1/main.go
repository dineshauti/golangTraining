package main

import "fmt"

func main() {

	pr1,pr2,pr3,pr4 := 1.0,1.0,1.0,1.0

	for i:=0; i<1000; i++ {

		pr1 = 0.15+0.85*(pr3/2)
		pr2 = 0.15+0.85*(pr1/2 + pr3/2)
		pr3 = 0.15+0.85*(pr4)
		pr4 = 0.15+0.85*(pr1/2 + pr2)
	}

	fmt.Println(pr1)
	fmt.Println(pr2)
	fmt.Println(pr3)
	fmt.Println(pr4)

}
