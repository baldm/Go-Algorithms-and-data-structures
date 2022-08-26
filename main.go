package main

import "Go-Algorithms-and-data-structures/data-structures"

func main() {
	// Used to call the different assignments

	var nodeAmount, edgesMatrix = QuickGraphInputParser()

	graph := data_structures.ListGraphConstructor(nodeAmount, edgesMatrix)
	
	return
}
