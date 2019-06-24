import "sort"

func minIncrementForUnique(A []int) int {
	o := 0
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		a, ap := A[i], A[i-1]
		if a <= ap {
			o += ap - a + 1
			A[i] = ap + 1
		}
	}
	return o
}
