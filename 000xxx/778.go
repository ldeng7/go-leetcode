import "math"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

var dirs [4][2]int = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) int {
	l := len(grid)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
		for j := 0; j < l; j++ {
			t[i][j] = math.MaxInt64
		}
	}

	var cal func(y, x, g int)
	cal = func(y, x, g int) {
		if y < 0 || y >= l || x < 0 || x >= l || max(g, grid[y][x]) >= t[y][x] {
			return
		}
		t[y][x] = max(g, grid[y][x])
		for _, d := range dirs {
			cal(y+d[0], x+d[1], t[y][x])
		}
	}

	cal(0, 0, grid[0][0])
	return t[l-1][l-1]
}
