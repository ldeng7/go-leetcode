func numMagicSquaresInside(grid [][]int) int {
	o, m, n := 0, len(grid), len(grid[0])
	check := func(i, j int) bool {
		m := [10]bool{}
		for h := i; h < i+2; h++ {
			for k := j; k < j+2; k++ {
				v := grid[h][k]
				if v < 1 || v > 9 || m[v] {
					return false
				}
				m[v] = true
			}
		}
		for k := 0; k <= 2; k++ {
			if 15 != grid[i+k][j]+grid[i+k][j+1]+grid[i+k][j+2] {
				return false
			}
		}
		for k := 0; k <= 2; k++ {
			if 15 != grid[i][j+k]+grid[i+1][j+k]+grid[i+2][j+k] {
				return false
			}
		}
		if 15 != grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] {
			return false
		}
		if 15 != grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] {
			return false
		}
		return true
	}
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			if grid[i+1][j+1] == 5 && check(i, j) {
				o++
			}
		}
	}
	return o
}
