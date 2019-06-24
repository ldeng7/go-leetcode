import "math"

func pacificAtlantic(matrix [][]int) [][]int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return [][]int{}
	}
	o, m, n := [][]int{}, len(matrix), len(matrix[0])
	p, a := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i], a[i] = make([]bool, n), make([]bool, n)
	}
	var cal func([][]bool, int, int, int)
	cal = func(v [][]bool, pre, i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || v[i][j] || matrix[i][j] < pre {
			return
		}
		pre, v[i][j] = matrix[i][j], true
		cal(v, pre, i+1, j)
		cal(v, pre, i-1, j)
		cal(v, pre, i, j+1)
		cal(v, pre, i, j-1)
	}
	for i := 0; i < m; i++ {
		cal(p, math.MinInt64, i, 0)
		cal(a, math.MinInt64, i, n-1)
	}
	for j := 0; j < n; j++ {
		cal(p, math.MinInt64, 0, j)
		cal(a, math.MinInt64, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				o = append(o, []int{i, j})
			}
		}
	}
	return o
}
