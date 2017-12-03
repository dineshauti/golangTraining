package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

var(
	lines []string
	V int
	err error
	data []string
	comp componentsDS
)

type componentsDS struct {
	connected bool
	numberOfComponents int
	components [][]int
	compMaps []map[int]int
}


// Helper function to check errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readGraph(path string) []string {

	file_path := path
	// Open file and create scanner on top of it
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func adjMatrix(lines []string) [][]float64 {

	V,err = strconv.Atoi(lines[0])
	check(err)

	adj_matrix := make([][]float64, V)
	for i := range adj_matrix {
		adj_matrix[i] = make([]float64, V)
	}

	for count:=2; count<len(lines); count++ {
		s := strings.Split(lines[count],",")
		u, _ := strconv.Atoi(s[0])
		v,_ := strconv.Atoi(s[1])

		adj_matrix[u-1][v-1],err = strconv.ParseFloat(s[2],64)
		adj_matrix[v-1][u-1],err = strconv.ParseFloat(s[2],64)
		check(err)
	}

	return adj_matrix
}

func adjList(lines []string) [][]int {

	V,err = strconv.Atoi(lines[0])
	check(err)

	vertices := make([][]int, V)
	for i := range vertices {
		vertices[i] = make([]int, 0, V)
	}

	for count:=2; count<len(lines); count++ {
		s := strings.Split(lines[count],",")
		u,_ := strconv.Atoi(s[0])
		v,_ := strconv.Atoi(s[1])

		vertices[u-1] = append(vertices[u-1], v)
		vertices[v-1] = append(vertices[v-1], u)
	}

	return vertices

}

func minKey(key []float64, mstSet []bool, vertices int) int {

	min := 9999.0
	minIndex := -1

	for v:=0; v<vertices; v++ {
		if !mstSet[v] && key[v] <= min {
			min = key[v];
			minIndex = v;
		}
	}

	return minIndex
}

func printMST(parent []int, vertices int, graph [][]float64, k int) {

	fmt.Println("Edge Weight")


	for i :=1; i < vertices; i++ {
		fmt.Printf("%d - %d    %f \n", comp.compMaps[k][parent[i]+1], comp.compMaps[k][i+1], graph[i][parent[i]])
	}




}

func prims(adj_matrix [][]float64, vertices int, k int) {

	parent := make([]int, vertices) // Array to store constructed MST
	key := make([]float64, vertices) // Key values used to pick minimum weight edge in cut
	mstSet := make([]bool, vertices) 	// To represent set of vertices not yet included in MST

	// Initialize all keys as INFINITE
	for i:=0; i<vertices; i++ {
		key[i] = 9999.0
	}

	// Always include first 1st vertex in MST.
	// Make key 0 so that this vertex is picked as first vertex
	key[0] = 0.0;

	// First node is always root of MST
	parent[0] = -1;

	for count:=0; count<vertices-1; count++ {

		u := minKey(key, mstSet, vertices);

		// Add the picked vertex to the MST Set
		mstSet[u] = true;

		// Update key value and parent index of the adjacent
		// vertices of the picked vertex. Consider only those
		// vertices which are not yet included in MST
		for v := 0; v < vertices; v++ {

			// graph[u][v] is non zero only for adjacent vertices of m
			// mstSet[v] is false for vertices not yet included in MST
			// Update the key only if graph[u][v] is smaller than key[v]
			if adj_matrix[u][v] != 0.0 && !mstSet[v] && adj_matrix[u][v] <  key[v] {

				parent[v]  = u;
				key[v] = adj_matrix[u][v];
			}
		}
	}

	// Print the constructed MST
	printMST(parent, vertices, adj_matrix, k);
}

func DFSUtil(v int, visited []bool, adj_list [][]int, conn map[int]int) {

	visited[v] = true

	conn[v] = v

	adj := adj_list[v-1]

	for i:=0; i<len(adj); i++ {
		if !visited[adj[i]] {
			DFSUtil(adj[i], visited, adj_list, conn)
		}
	}
}

func connectedComponents(adj_list [][]int) {

	components := make(map[int]int)
	visited := make([]bool, V+1)

	countConnSize:=0
	numComponents := 0

	comp.components = make([][]int, 3, V) // I have hard coded 3 as the number of components. As I am not able to determine it dynamically



	for i := range comp.components {
		comp.components[i] = make([]int, 0, V)
	}

	for v:=1; v<=V; v++ {

		if visited[v] == false {
			DFSUtil(v, visited, adj_list, components)

			if len(components) == V {

				for x := range components {
					comp.components[0] = append(comp.components[0], x)
				}
				comp.connected = true
				numComponents = 1

			} else {

				numComponents++

				//fmt.Printf("Component %d :- ", numComponents)
				for x := range components {
					//fmt.Printf("%d ",x)

					comp.components[numComponents-1] = append(comp.components[numComponents-1], x)

				}
				countConnSize++
				comp.connected = false
			}

			for x := range components {
				delete(components, x)
			}
		}
		comp.numberOfComponents = numComponents
	}

}

func init() {
	// Read the graph
	data = readGraph("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\golangTraining\\project_dataStructures\\read_file\\graph4.md")

}

func main() {



	// Setup Adjacency Matrix/List
	adj_matrix := adjMatrix(data)
	adj_list := adjList(data)

	// Print Adjacency Matrix/List
	fmt.Print("Adjacency Matrix: ")
	fmt.Println(adj_matrix)
	fmt.Print("Adjacency List: ")
	fmt.Println(adj_list)

	connectedComponents(adj_list)
	fmt.Println()

	if comp.connected {
		fmt.Println("Graph is connected")
	} else {
		fmt.Println("Graph is not connected")
	}

	fmt.Printf("Number of components :- %d \n",comp.numberOfComponents)

	for i:=0; i<comp.numberOfComponents; i++ {
		fmt.Printf("Component %d :- %d \n", i+1, comp.components[i])

	}

	comp.compMaps = make([]map[int]int, comp.numberOfComponents)
	for i := range comp.compMaps {
		comp.compMaps[i] = make(map[int]int, 10)
	}


	fmt.Println()

	if comp.numberOfComponents != 1 {

		for k:=0; k<comp.numberOfComponents; k++ {

			sort.Ints(comp.components[k])

			for i:=0; i<len(comp.components[k]); i++ {
				comp.compMaps[k][i+1] = comp.components[k][i]
			}

			// Form adjacency matrices for the components
			compMatrix := make([][]float64,len(comp.components[k]))
			for m := range compMatrix {
				compMatrix[m] = make([]float64,len(comp.components[k]))
			}

			for i:=0; i<len(comp.components[k]); i++ {
				x := comp.components[k][i]
				for j:=0; j<len(comp.components[k]); j++ {
					y := comp.components[k][j]
					compMatrix[i][j] = adj_matrix[x-1][y-1]
				}
			}

			fmt.Printf("Component %d :- \n", k+1)
			prims(compMatrix, len(comp.components[k]), k)

			fmt.Println()
		}

	} else {

		sort.Ints(comp.components[0])

		for i:=0; i<len(comp.components[0]); i++ {
			comp.compMaps[0][i+1] = comp.components[0][i]
		}

		prims(adj_matrix, V, 0)
	}


}
