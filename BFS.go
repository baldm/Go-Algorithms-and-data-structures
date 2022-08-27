package main

import "fmt"

func isInList(list []int, value int) bool {
	for i := range list {
		if list[i] == value {
			return true
		}
	}
	return false
}

func BFS(graph map[int][]int, startNode int, endNode int) int {
	// Finds the shortest path between nodes

	q := make([]int, 0)       // Queue for the BFS
	visited := make([]int, 0) // List of nodes we have visited
	path := 0                 // Length of the path
	shortestPath := 1844674407370955161

	q = append(q, startNode)

	for len(q) > 0 {
		// Pop the value from the list
		node := q[0]
		//fmt.Printf("Popping %d from queue \n", node)
		q = q[1:]
		//fmt.Printf("Queue is now %v \n", q)
		if isInList(visited, node) {
			continue
		} else {
			visited = append(visited, node)
			//fmt.Printf("Visited list is now  %v \n", visited)
		}
		children := graph[node]

		//fmt.Printf("Searching in children: %v \n", children)

		path++

		for i := range children {
			//fmt.Printf("Looking at child %d \n", children[i])

			// Check if we found the final node
			if children[i] == endNode {

				fmt.Printf("Found the end node at range %d \n", path)
				if path < shortestPath {
					shortestPath = path
				}
			}

			// Add to queue if the child has child nodes
			if len(graph[children[i]]) > 0 && !isInList(visited, children[i]) {
				//fmt.Printf("Adding  %d to the queue \n", children[i])
				q = append(q, children[i])
			} else {
				path--
			}

		}
		//fmt.Println()
	}

	return shortestPath
}
