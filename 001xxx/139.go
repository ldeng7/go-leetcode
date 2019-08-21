func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func largest1BorderedSquare(grid [][]int) int {
	o, m, n := 0, len(grid), len(grid[0])
	h, v := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		h[i], v[i] = make([]int, n), make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if 0 != j {
					h[i][j] = h[i][j-1] + 1
				} else {
					h[i][j] = 1
				}
				if 0 != i {
					v[i][j] = v[i-1][j] + 1
				} else {
					v[i][j] = 1
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for k := min(h[i][j], v[i][j]); k > o; k-- {
				if v[i][j-k+1] >= k && h[i-k+1][j] >= k {
					o = k
				}
			}
		}
	}
	return o * o
}
