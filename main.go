package main

import (
	"Go-Algorithms-and-data-structures/algorithms"
	"fmt"
)

func main() {
	// Used to call the different assignments

	var n = TwoLines()
	result := algorithms.Sort(n)

	for _, number := range result {
		fmt.Printf("%d ", number)
	}

	return
}
