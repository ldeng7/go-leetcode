import "sort"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func twoSumLessThanK(A []int, K int) int {
	sort.Ints(A)
	o, l, r := 0, 0, len(A)-1
	for l < r {
		s := A[l] + A[r]
		if s >= K {
			r--
		} else {
			o, l = max(o, s), l+1
		}
	}
	if 0 == o {
		return -1
	}
	return o
}
