package algorithms

func FindMinN(numbers []int, numbersToFind []int) []int {
	// Assignment: Find array of numbers in a bigger array of numbers
	// Recursively search and print as you find them
	var results []int
	for _, numberToFind := range numbersToFind {
		result := binarySearch(numbers, numberToFind, 0, len(numbers)-1)
		results = append(results, result)
	}
	return results
}

func binarySearch(numbers []int, val int, left int, right int) int {
	var mid int

	// Not found
	if right < left {
		return -1
	}

	mid = (left + right) / 2

	// This is the number
	if numbers[mid] == val {
		return mid
	} else if numbers[mid] < val {
		// Search right
		return binarySearch(numbers, val, mid+1, right)
	} else {
		// Search left
		return binarySearch(numbers, val, left, mid-1)
	}
}
