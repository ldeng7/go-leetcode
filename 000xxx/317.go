import "math"

func copyGrid(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	out := make([][]int, m)
	for i := 0; i < m; i++ {
		out[i] = make([]int, n)
		copy(out[i], grid[i])
	}
	return out
}

func shortestDistance(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return -1
	}
	out, v := math.MaxInt64, 0
	m, n := len(grid), len(grid[0])
	sum := copyGrid(grid)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			out = math.MaxInt64
			dist := copyGrid(grid)
			qi, qj := []int{i}, []int{j}
			for 0 != len(qi) {
				ti, tj := qi[0], qj[0]
				qi, qj = qi[1:], qj[1:]
				for _, dir := range dirs {
					y, x := ti+dir[0], tj+dir[1]
					if y >= 0 && y < m && x >= 0 && x < n && grid[y][x] == v {
						grid[y][x]--
						dist[y][x] = dist[ti][tj] + 1
						sum[y][x] += dist[y][x] - 1
						qi, qj = append(qi, y), append(qj, x)
						if sum[y][x] < out {
							out = sum[y][x]
						}
					}
				}
			}
			v--
		}
	}
	if out == math.MaxInt64 {
		return -1
	}
	return out
}
