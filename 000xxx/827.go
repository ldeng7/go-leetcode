func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	o, k, ar := 0, 2, []int{}
	var cal func(int, int) int
	cal = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = k
		return 1 + cal(i-1, j) + cal(i+1, j) + cal(i, j-1) + cal(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			l := cal(i, j)
			o, ar, k = max(o, l), append(ar, l), k+1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			s := map[int]bool{}
			if i > 0 {
				s[grid[i-1][j]] = true
			}
			if j > 0 {
				s[grid[i][j-1]] = true
			}
			if i < m-1 {
				s[grid[i+1][j]] = true
			}
			if j < n-1 {
				s[grid[i][j+1]] = true
			}
			l := 0
			for k, _ = range s {
				if k != 0 {
					l += ar[k-2]
				}
			}
			o = max(o, l+1)
		}
	}
	return o
}
