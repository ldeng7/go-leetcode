func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	o := make([]int, n)
lj:
	for j := 0; j < n; j++ {
		k := j
		for i := 0; i < m; i++ {
			t := grid[i][k]
			k += t
			if k >= n || k < 0 || grid[i][k]*t == -1 {
				o[j] = -1
				continue lj
			}
		}
		o[j] = k
	}
	return o
}
