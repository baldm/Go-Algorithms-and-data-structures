package main

import (
	"Go-Algorithms-and-data-structures/algorithms"
	"fmt"
)

func main() {
	// Used to call the different assignments

	var edgesMatrix = WeightedGraphInputParser()

	fmt.Println(algorithms.Kruskals(len(edgesMatrix)+1, edgesMatrix))

	return
}
