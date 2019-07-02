func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func minDeletionSize(A []string) int {
	l, l0 := len(A), len(A[0])
	ar := make([]int, l0)
	for i := 0; i < l0; i++ {
		ar[i] = 1
	}
	m := 1
	for i := 1; i < l0; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k <= l; k++ {
				if k == l {
					ar[i] = max(ar[i], ar[j]+1)
					m = max(m, ar[i])
				} else if A[k][j] > A[k][i] {
					break
				}
			}
		}
	}
	return l0 - m
}
