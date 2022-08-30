package algorithms

import "fmt"

func DFS(graph map[int][]int, startNode int) {
	// Print the path the DFS takes through the graph

	var visited = make(map[int]bool) // Path length to all nodes in the graph

	// Initialize our visited nodes
	visited[startNode] = true

	// Begin search
	DFSSearch(graph, startNode, visited)

	// Final print leaving the graph
	fmt.Println(startNode)
	return
}

func DFSSearch(graph map[int][]int, node int, visited map[int]bool) {
	// Recursively Depth-first searches an incidence graph
	// Prints the first and the last time an element is visited

	// Mark as visited and print
	visited[node] = true
	fmt.Println(node) // First time visited

	// Get children and loop over them
	children := graph[node]
	for i := range children {
		child := children[i]

		// If the child has not been visited
		if !visited[child] {
			DFSSearch(graph, child, visited) // Recursive search
			fmt.Println(child)               // Last time visited
		}

	}

}
