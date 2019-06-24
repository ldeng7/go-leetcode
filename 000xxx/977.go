import "sort"

func sortedSquares(A []int) []int {
	for i, a := range A {
		A[i] = a * a
	}
	sort.Ints(A)
	return A
}
