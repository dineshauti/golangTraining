package main

import "fmt"

var(
	vertices int
)


func minKey(key []int, mstSet []bool) int {

	min := 9999
	minIndex := -1

	for v:=0; v<vertices; v++ {
		if !mstSet[v] && key[v] <= min {
		min = key[v]
		minIndex = v
		}
	}

	return minIndex
}

func printMST(parent []int, vertices int, graph [][]int) {

	fmt.Println("Edge Weight")

	for i :=1; i < vertices; i++ {
		fmt.Printf("%d - %d    %d \n", parent[i]+1, i+1, graph[i][parent[i]])
	}

}

func prims(graph [][]int) {

	vertices = len(graph)

	// Array to store constructed MST
	parent := make([]int, vertices)

	// Key values used to pick minimum weight edge in cut
	key := make([]int, vertices)

	// To represent set of vertices not yet included in MST
	mstSet := make([]bool, vertices)

	// Initialize all keys as INFINITE
	for i:=0; i<vertices; i++ {
		key[i] = 9999
	}

	// Always include first 1st vertex in MST.
	// Make key 0 so that this vertex is picked as first vertex
	key[0] = 0

	// First node is always root of MST
	parent[0] = -1

	for count:=0; count<vertices-1; count++ {

		u := minKey(key, mstSet)

		// Add the picked vertex to the MST Set
		mstSet[u] = true


		// Update key value and parent index of the adjacent
		// vertices of the picked vertex. Consider only those
		// vertices which are not yet included in MST
		for v := 0; v < vertices; v++ {

			// graph[u][v] is non zero only for adjacent vertices of m
			// mstSet[v] is false for vertices not yet included in MST
			// Update the key only if graph[u][v] is smaller than key[v]
			if graph[u][v] != 0 && !mstSet[v] && graph[u][v] <  key[v] {

			parent[v]  = u
			key[v] = graph[u][v]
			}
		}
	}


	// Print the constructed MST
	printMST(parent, vertices, graph);


}

func main() {

	graph := [][]int{{0,1,3},
					{1,0,2},
					{3,2,0}}



	prims(graph)


}
