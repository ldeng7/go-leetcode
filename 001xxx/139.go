func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	h, v := [101][101]int{}, [101][101]int{}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				h[i][j] = h[i][j-1] + 1
				v[i][j] = v[i-1][j] + 1
			}
		}
	}

	for k := min(m, n); k > 0; k-- {
		for i := 1; i+k-1 <= m; i++ {
			for j := 1; j+k-1 <= n; j++ {
				if v[i+k-1][j] >= k && v[i+k-1][j+k-1] >= k && h[i][j+k-1] >= k && h[i+k-1][j+k-1] >= k {
					return k * k
				}
			}
		}
	}
	return 0
}
