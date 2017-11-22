package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Age int `json:"wisdom score"`
}



func main() {
	p1 := person{
		"Dinesh",
		"Auti",
		27,
	}

	bs,_ := json.Marshal(p1)
	fmt.Println(string(bs))
	fmt.Printf("%T \n", bs)
}
