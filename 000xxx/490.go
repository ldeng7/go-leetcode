var dirs [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func hasPath(maze [][]int, start []int, destination []int) bool {
	if 0 == len(maze) || 0 == len(maze[0]) {
		return false
	}
	m, n := len(maze), len(maze[0])
	t := make([][]int8, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int8, n)
	}

	var cal func(i, j int) bool
	cal = func(i, j int) bool {
		if i == destination[0] && j == destination[1] {
			return true
		}
		if t[i][j] != 0 {
			return t[i][j] == 1
		}
		var b bool
		maze[i][j] = 2
		for _, d := range dirs {
			y, x := i, j
			for y >= 0 && y < m && x >= 0 && x < n && maze[y][x] != 1 {
				y, x = y+d[0], x+d[1]
			}
			y, x = y-d[0], x-d[1]
			if maze[y][x] != 2 {
				b = b || cal(y, x)
			}
		}
		if b {
			t[i][j] = 1
		} else {
			t[i][j] = -1
		}
		return b
	}
	return cal(start[0], start[1])
}
