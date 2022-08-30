package main

import (
	"Go-Algorithms-and-data-structures/algorithms"
	"Go-Algorithms-and-data-structures/data-structures"
	"fmt"
)

func main() {
	// Used to call the different assignments

	var nodeAmount, edgesMatrix = QuickGraphInputParser()

	graph := data_structures.ListGraphConstructor(edgesMatrix, true)

	fmt.Println(algorithms.BFS(graph, nodeAmount, 0, 1))

	return
}
