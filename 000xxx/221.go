func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maximalSquare(matrix [][]byte) int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	o, p := 0, 0
	t := make([]int, m+1)
	for j := 0; j < n; j++ {
		for i := 1; i <= m; i++ {
			ti := t[i]
			if matrix[i-1][j] == '1' {
				t[i] = min(t[i], min(t[i-1], p)) + 1
				o = max(o, t[i])
			} else {
				t[i] = 0
			}
			p = ti
		}
	}
	return o * o
}
