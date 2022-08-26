package main

import "Go-Algorithms-and-data-structures/data-structures"

func main() {
	// Used to call the different assignments

	var nodeAmount, edgesMatrix, Q1List, Q2List = QuickGraphInputParser()

	graph := data_structures.ListGraphConstructor(nodeAmount, edgesMatrix)

	for i := range Q1List {
		data_structures.ListIsNeighbors(Q1List[i][0], Q1List[i][1], graph)
	}

	for i := range Q2List {
		data_structures.ListPrintNeighbors(Q2List[i], graph)
	}

	return
}
