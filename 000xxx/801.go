import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSwap(A []int, B []int) int {
	n1, s1 := 0, 1
	for i := 1; i < len(A); i++ {
		n2, s2 := math.MaxInt64, math.MaxInt64
		if A[i-1] < A[i] && B[i-1] < B[i] {
			n2, s2 = min(n2, n1), min(s2, s1+1)
		}
		if A[i-1] < B[i] && B[i-1] < A[i] {
			n2, s2 = min(n2, s1), min(s2, n1+1)
		}
		n1, s1 = n2, s2
	}
	return min(n1, s1)
}
