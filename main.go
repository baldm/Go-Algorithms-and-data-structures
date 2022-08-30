package main

import (
	"Go-Algorithms-and-data-structures/algorithms"
	"Go-Algorithms-and-data-structures/data-structures"
)

func main() {
	// Used to call the different assignments

	var edgesMatrix = QuickGraphInputParser()

	graph := data_structures.ListGraphConstructor(edgesMatrix, true)

	algorithms.DFS(graph, 0)

	return
}
