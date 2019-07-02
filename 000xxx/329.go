var dirs = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestIncreasingPath(matrix [][]int) int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return 0
	}
	o, m, n := 0, len(matrix), len(matrix[0])
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
	}

	var cal func(int, int) int
	cal = func(i, j int) int {
		if 0 != t[i][j] {
			return t[i][j]
		}
		o := 1
		for _, d := range dirs {
			y, x := i+d[0], j+d[1]
			if y < 0 || y >= m || x < 0 || x >= n || matrix[y][x] <= matrix[i][j] {
				continue
			}
			o = max(o, cal(y, x)+1)
		}
		t[i][j] = o
		return o
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			o = max(o, cal(i, j))
		}
	}
	return o
}
