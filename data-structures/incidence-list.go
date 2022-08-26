package data_structures

import (
	"fmt"
)

func ListGraphConstructor(nodeAmount int, edgesMatrix [][]int) map[int][]int {

	// Build graph as map
	var graph = make(map[int][]int)

	// Populate with graph with edges
	for i := 0; i < len(edgesMatrix); i++ {
		c1 := edgesMatrix[i][0]
		c2 := edgesMatrix[i][1]
		graph[c1] = append(graph[c1], c2)
		graph[c2] = append(graph[c2], c1)
	}

	return graph
}

func ListIsNeighbors(firstCord int, secondCord int, graph map[int][]int) {
	// Prints yes/no if 2 coordinates are/are not neighbors
	for i := range graph[firstCord] {
		if graph[firstCord][i] == secondCord {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

func ListPrintNeighbors(value int, graph map[int][]int) {
	// Prints all neighbors of a value
	for i := range graph[value] {
		fmt.Printf("%d ", graph[value][i])
	}
	fmt.Println()
}
