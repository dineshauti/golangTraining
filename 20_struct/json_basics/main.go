package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type person struct {
	Craft string
	Name string
}

type space struct {
	Number int `json:"number"`
	People []person
}

func main() {
	res, err := http.Get("http://api.open-notify.org/astros.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP: %s\n", res.Status)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}


	p1 := space{}

	err2 := json.Unmarshal(body, &p1)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(p1.People)
	fmt.Println(p1.Number)

}
