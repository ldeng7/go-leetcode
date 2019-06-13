func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func cherryPickup(grid [][]int) int {
	l := len(grid)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
		for j := 0; j < l; j++ {
			t[i][j] = -1
		}
	}
	t[0][0] = grid[0][0]
	for i := 1; i < l<<1-1; i++ {
		for j := l - 1; j >= 0; j-- {
			for k := l - 1; k >= 0; k-- {
				if i-j < 0 || i-j >= l || i-k < 0 || i-k >= l || grid[j][i-j] < 0 || grid[k][i-k] < 0 {
					t[j][k] = -1
					continue
				}
				if j > 0 && t[j-1][k] > t[j][k] {
					t[j][k] = t[j-1][k]
				}
				if k > 0 && t[j][k-1] > t[j][k] {
					t[j][k] = t[j][k-1]
				}
				if j > 0 && k > 0 && t[j-1][k-1] > t[j][k] {
					t[j][k] = t[j-1][k-1]
				}
				if t[j][k] >= 0 {
					t[j][k] += grid[j][i-j]
					if j != k {
						t[j][k] += grid[k][i-k]
					}
				}
			}
		}
	}
	if t[l-1][l-1] > 0 {
		return t[l-1][l-1]
	}
	return 0
}
