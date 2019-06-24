func transpose(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	o := make([][]int, n)
	for j := 0; j < n; j++ {
		o[j] = make([]int, m)
		for i := 0; i < m; i++ {
			o[j][i] = A[i][j]
		}
	}
	return o
}
