package data_structures

import "fmt"

func MatrixGraphConstructor(nodeAmount int, edgesMatrix [][]int) [][]int {
	// Build graph as matrix
	graph := make([][]int, nodeAmount)

	// Populate with graph with edges
	for i := range graph {
		graph[i] = make([]int, nodeAmount)
	}

	// Populate with edges
	for i := range edgesMatrix {
		firstCord := edgesMatrix[i][0]
		secondCord := edgesMatrix[i][1]
		graph[firstCord][secondCord] = 1
		graph[secondCord][firstCord] = 1
	}
	return graph
}

func MatrixIsNeighbors(firstCord int, secondCord int, graph [][]int) {
	// Prints yes/no if 2 coordinates are/are not neighbors
	if graph[firstCord][secondCord] == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func MatrixPrintNeighbors(value int, graph [][]int) {
	// Prints all neighbors of a value
	for i := 0; i < len(graph); i++ {
		if graph[value][i] == 1 {
			fmt.Printf("%d ", i)
		}
	}
}
