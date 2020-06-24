const N = 70

var t [N][N][N]int

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var cal func(int, int, int) int
	cal = func(i, j, k int) int {
		if i >= m || j > k {
			return 0
		}
		p := &(t[i][j][k])
		v := *p
		if v != -1 {
			return v
		}

		v = grid[i][j]
		if j != k {
			v += grid[i][k]
		}
		ma := 0
		for dj := -1; dj <= 1; dj++ {
			if j+dj < 0 || j+dj >= n {
				continue
			}
			for dk := -1; dk <= 1; dk++ {
				if k+dk < 0 || k+dk >= n {
					continue
				}
				ma = max(ma, cal(i+1, j+dj, k+dk))
			}
		}
		v += ma
		*p = v
		return v
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ar := &(t[i][j])
			for k := 0; k < n; k++ {
				(*ar)[k] = -1
			}
		}
	}
	return cal(0, 0, n-1)
}
