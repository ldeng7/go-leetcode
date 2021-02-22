func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	var l, r, mi int = 0, 2e6, 0
	var v [][]bool
	var cal func(int, int) bool
	cal = func(i, j int) bool {
		if i == m-1 && j == n-1 {
			return true
		}
		v[i][j] = true
		h := heights[i][j]
		return (i > 0 && !v[i-1][j] && abs(h-heights[i-1][j]) <= mi && cal(i-1, j)) ||
			(j > 0 && !v[i][j-1] && abs(h-heights[i][j-1]) <= mi && cal(i, j-1)) ||
			(i < m-1 && !v[i+1][j] && abs(h-heights[i+1][j]) <= mi && cal(i+1, j)) ||
			(j < n-1 && !v[i][j+1] && abs(h-heights[i][j+1]) <= mi && cal(i, j+1))
	}

	for l < r {
		mi = l + (r-l)>>1
		v = make([][]bool, m)
		for i := 0; i < m; i++ {
			v[i] = make([]bool, n)
		}
		if cal(0, 0) {
			r = mi
		} else {
			l = mi + 1
		}
	}
	return l
}
