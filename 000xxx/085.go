func maximalRectangle(matrix [][]byte) int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return 0
	}
	out := 0
	m, n := len(matrix), len(matrix[0])
	h, l, r := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n
	}
	for i := 0; i < m; i++ {
		il, ir := 0, n
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				h[j]++
				if il > l[j] {
					l[j] = il
				}
			} else {
				h[j], l[j], il = 0, 0, j+1
			}
		}
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if ir < r[j] {
					r[j] = ir
				}
			} else {
				r[j], ir = n, j
			}
			o := (r[j] - l[j]) * h[j]
			if o > out {
				out = o
			}
		}
	}
	return out
}
