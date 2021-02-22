func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	o := make([][]int, m)
	for i := 0; i < m; i++ {
		o[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r, c := rowSum[i], colSum[j]
			m := min(r, c)
			o[i][j] = m
			rowSum[i], colSum[j] = r-m, c-m
		}
	}
	return o
}
