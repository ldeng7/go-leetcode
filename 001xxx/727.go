import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	for j := 0; j < n; j++ {
		for i, p := 0, 0; i < m; i++ {
			v := matrix[i][j]
			p = p*v + v
			matrix[i][j] = p
		}
	}
	o := 0
	for _, r := range matrix {
		sort.Ints(r)
		for i, v := range r {
			o = max(o, v*(n-i))
		}
	}
	return o
}
