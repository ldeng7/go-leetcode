func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	l, r := matrix[0][0], matrix[m-1][n-1]
	for l < r {
		h := l + (r-l)>>1
		i, j, c := m-1, 0, 0
		for i >= 0 && j < m {
			if matrix[i][j] <= h {
				c, j = c+i+1, j+1
			} else {
				i--
			}
		}
		if c < k {
			l = h + 1
		} else {
			r = h
		}
	}
	return l
}
