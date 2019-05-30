func maxKilledEnemies(grid [][]byte) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	out, cr := 0, 0
	cc := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 || grid[i][j-1] == 'W' {
				cr = 0
				for k := j; k < n && grid[i][k] != 'W'; k++ {
					if grid[i][k] == 'E' {
						cr++
					}
				}
			}
			if i == 0 || grid[i-1][j] == 'W' {
				cc[j] = 0
				for k := i; k < m && grid[k][j] != 'W'; k++ {
					if grid[k][j] == 'E' {
						cc[j]++
					}
				}
			}
			if grid[i][j] == '0' && cr+cc[j] > out {
				out = cr + cc[j]
			}
		}
	}
	return out
}
