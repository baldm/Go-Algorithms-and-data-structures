package algorithms

// FindPeak1 Finds the biggest number in slice
// with smaller numbers on either side
func FindPeak1(n int, A []int) int {

	if A[0] >= A[1] {
		return 0
	}

	for i := 1; i <= n-2; i++ {
		if A[i-1] <= A[i] && A[i] >= A[i+1] {
			return i
		}
	}
	if A[n-2] <= A[n-1] {
		return n - 1
	}
	return 0
}

// FindPeak2 Recursively search for peak in slice
func FindPeak2(A []int, i int, j int) int {
	var m int
	m = (i + j) / 2

	if (m == 0 || A[m] >= A[m-1]) && (m == len(A)-1 || A[m] >= A[m+1]) {
		return m
	} else if m > 0 && A[m-1] > A[m] {
		return FindPeak2(A, i, m-1)
	} else if m < len(A)-1 && A[m] < A[m+1] {
		return FindPeak2(A, m+1, j)
	}
	return -1

}
