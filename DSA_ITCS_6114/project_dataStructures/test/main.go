package main

import (
	"fmt"
)


func main() {

	adj_forest := make([][]float64, 3)
	for i := range adj_forest {
		adj_forest[i] = make([]float64, 3)
	}

	adj_forest[0][2] = 32

	fmt.Println(adj_forest)


}
