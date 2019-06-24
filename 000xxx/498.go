func findDiagonalOrder(matrix [][]int) []int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	l := m * n
	r, c, o := 0, 0, make([]int, l)
	for i := 0; i < l; i++ {
		o[i] = matrix[r][c]
		if (r+c)&1 == 0 {
			if c == n-1 {
				r++
			} else if r == 0 {
				c++
			} else {
				r, c = r-1, c+1
			}
		} else {
			if r == m-1 {
				c++
			} else if c == 0 {
				r++
			} else {
				r, c = r+1, c-1
			}
		}
	}
	return o
}
