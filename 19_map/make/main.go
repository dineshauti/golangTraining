package main

import "fmt"

func main() {

	// Way one to innitalize a map
	myGreetings := make(map[string]string)
	myGreetings["Dinesh"] = "Good Morning"
	myGreetings["Neha"] = "Namaste"

	// Way one to innitalize a map
/*	myGreetings1 := map[string]string{
		"Dinesh": "Good Morning",
		"Neha": "Namaste",
	}*/

	// comma, ok idiom
	if v, ok := myGreetings["Dinesh"]; ok {
		fmt.Println(v)
	}


}
