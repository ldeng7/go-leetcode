func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	o, l := 0, len(A)
	a, b := make([]int, l), make([]int, l)
	for i := 1; i < l; i++ {
		A[i] += A[i-1]
	}
	for i := 0; i < l; i++ {
		if i < L {
			a[i] = A[i]
		} else {
			a[i] = max(a[i-1], A[i]-A[i-L])
		}
		if i < M {
			b[i] = A[i]
		} else {
			b[i] = max(b[i-1], A[i]-A[i-M])
		}
		if i < L+M {
			o = A[i]
		} else {
			o = max(o, max(A[i]-A[i-L]+b[i-L], A[i]-A[i-M]+a[i-M]))
		}
	}
	return o
}
