import "math"

var dirs = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	mi := math.MaxInt64
	grid[0][0] = 2
	var cal func(int, int, int, int) int
	cal = func(i, j, r, c int) int {
		if r < 0 {
			return math.MaxInt64
		} else if mi == m+n-2 || c >= mi {
			return mi
		} else if i == m-1 && j == n-1 {
			mi = min(mi, c+1)
			return c + 1
		}
		o := math.MaxInt64
		for _, d := range dirs {
			i1, j1 := i+d[0], j+d[1]
			if i1 < 0 || i1 >= m || j1 < 0 || j1 >= n {
				continue
			}
			if x := grid[i1][j1]; x != 2 {
				grid[i1][j1] = 2
				o = min(o, cal(i1, j1, r-x, c+1))
				grid[i1][j1] = x
			}
		}
		return o
	}

	if o := cal(0, 0, k, -1); o != math.MaxInt64 {
		return o
	}
	return -1
}
