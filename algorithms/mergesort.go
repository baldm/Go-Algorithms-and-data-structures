package algorithms

func Sort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	// Divide list
	A1 := numbers[0 : (len(numbers)+1)/2]
	A2 := numbers[(len(numbers)+1)/2:]

	// Recursively sort the two new lists
	A1 = Sort(A1)
	A2 = Sort(A2)

	// Merge the sorted lists
	return merge(A1, A2)
}

func merge(A1 []int, A2 []int) []int {
	p, p1, p2 := 0, 0, 0
	A := make([]int, len(A1)+len(A2))

	// Go over each pile and simultaneously
	// merge into the bigger list
	for p1 < len(A1) && p2 < len(A2) {
		if A1[p1] < A2[p2] {
			A[p] = A1[p1]
			p1++
		} else {
			A[p] = A2[p2]
			p2++
		}
		p++
	}

	// Go over the rest that was missed
	// in the above loop
	for p1 < len(A1) {
		A[p] = A1[p1]
		p1++
		p++
	}
	for p2 < len(A2) {
		A[p] = A2[p2]
		p2++
		p++
	}

	return A
}
