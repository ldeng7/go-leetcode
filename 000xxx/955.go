func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	o, t := 0, make([]bool, m)
	for j := 0; j < n; j++ {
		t1 := make([]bool, m)
		i, c := 1, 1
		for ; i < m; i++ {
			if t[i] || A[i][j] > A[i-1][j] {
				t1[i], c = true, c+1
			} else if A[i-1][j] > A[i][j] {
				o++
				break
			}
		}
		if c == m {
			break
		}
		if i == m {
			t = t1
		}
	}
	return o
}
