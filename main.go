package main

import (
	"fmt"
)

func main() {
	// Used to call the different assignments

	var edgesMatrix = QuickGraphInputParser()

	graph := ListGraphConstructor(edgesMatrix, true)

	fmt.Println(BFS(graph, 0, 1))

	return
}
