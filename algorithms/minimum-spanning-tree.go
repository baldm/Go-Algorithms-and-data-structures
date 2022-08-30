package algorithms

func Kruskals(nodeAmount int, graph [][]int) int {
	// Calculates the minimum weight of the spanning tree

	miniumWeight := 0
	nodeList := incrementingSlice(nodeAmount) // Goes [0, 1, 2,...]
	sz := repeatedSlice(1, nodeAmount)        // Goes [1, 1, 1,...]

	// Go through all the nodes
	for _, element := range graph {

		// Find the x and y cords
		find1 := find(element[0], nodeList)
		find2 := find(element[1], nodeList)

		// If they are not the same
		if find1 != find2 {
			// Add to the sum of the weights
			miniumWeight += element[2]
			// Group the nodes
			nodeList, sz = union(find1, find2, nodeList, sz)

		}

	}

	return miniumWeight
}

func find(i int, nodeList []int) int {
	// Find the group which the value belongs too
	for i != nodeList[i] {
		i = nodeList[i]
	}
	return nodeList[i]
}

func union(i int, j int, nodeList []int, sz []int) ([]int, []int) {
	// Find the groups of the nodes
	rI := find(i, nodeList)
	rJ := find(j, nodeList)

	// Collect the groups into one major group
	if rI != rJ {
		if sz[rI] < sz[rJ] {
			nodeList[rI] = rJ
			sz[rJ] += sz[rI]
		} else {
			nodeList[rJ] = rI
			sz[rI] += sz[rJ]
		}
	}
	return nodeList, sz
}

// Sourced from @squiguy
// at https://stackoverflow.com/questions/62745376/how-to-create-a-slice-with-repeated-elements
func repeatedSlice(value, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = value
	}
	return arr
}

func incrementingSlice(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}
