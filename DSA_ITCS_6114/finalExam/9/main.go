package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

var(
	V int
	lines      []string
	err        error
	data       []string
	comp componentsDS
)

type componentsDS struct {
	connected bool
}

// Helper function to check errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Helper function to read graphs from the input file
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

// Takes input the file read and outputs Adjacency list
func adjList(lines []string) [][]int {

	V, err = strconv.Atoi(lines[0])
	check(err)

	vertices := make([][]int, V)
	for i := range vertices {
		vertices[i] = make([]int, 0, V)
	}

	for count := 2; count < len(lines); count++ {
		s := strings.Split(lines[count], ",")
		u, _ := strconv.Atoi(s[0])
		v, _ := strconv.Atoi(s[1])

		vertices[u-1] = append(vertices[u-1], v)
		vertices[v-1] = append(vertices[v-1], u)
	}

	return vertices

}

func dfSUtil(v int, visited []bool, adj_list [][]int, conn map[int]int) {

	visited[v] = true
	conn[v] = v

	adj := adj_list[v-1]

	for i := 0; i < len(adj); i++ {
		if !visited[adj[i]] {
			dfSUtil(adj[i], visited, adj_list, conn)
		}
	}
}

func connectedComponents(adj_list [][]int) {

	components := make(map[int]int)
	visited := make([]bool, V+1)

	for v := 1; v <= V; v++ {

		if visited[v] == false {
			dfSUtil(v, visited, adj_list, components)

			if len(components) == V {
				comp.connected = true
			} else {
				comp.connected = false
			}

			for x := range components {
				delete(components, x)
			}
		}
	}
}

func init() {
	// Read the graph
	data = readGraph("C:\\Users\\Dinesh Auti\\IdeaProjects\\go\\src\\github.com\\dineshauti\\golangTraining\\DSA_ITCS_6114\\project_dataStructures\\read_file\\graph2.md")

}

func main() {

	adj_list := adjList(data)
	connectedComponents(adj_list)

	if comp.connected {
		fmt.Println("Graph is connected")
	} else {
		fmt.Println("Graph is not connected")
	}
}
