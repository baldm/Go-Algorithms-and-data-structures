package algorithms

func BFS(graph map[int][]int, startNode int, endNode int) int {
	// Finds the shortest path between nodes

	q := make([]int, 0)               // Queue for the BFS
	var distances = make(map[int]int) // Path length to all nodes in the graph

	// Initialize our queue and path length
	q = append(q, startNode)
	distances[startNode] = 0

	// Go as long as the queue has elements
	for len(q) > 0 {

		// Get the node from the queue and pop the element
		node := q[0]
		q = q[1:]

		// Get the current distance for the algorithm
		distance := distances[node]

		// Get the children of the node
		children := graph[node]

		// Go over all children
		for i := range children {
			child := children[i]

			// Do we already have the child path
			_, doesMapContains := distances[child]

			// If not we add the child to the queue and register the path
			if !doesMapContains {
				q = append(q, child)
				distances[child] = distance + 1
			}
		}
	}

	return distances[endNode]
}
