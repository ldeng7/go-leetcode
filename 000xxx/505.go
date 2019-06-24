import "math"

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = math.MaxInt64
		}
	}
	t[start[0]][start[1]] = 0
	q := [][2]int{{start[0], start[1]}}
	for 0 != len(q) {
		h := q[0]
		q = q[1:]
		for _, dir := range dirs {
			y, x := h[0], h[1]
			dist := t[y][x]
			for y >= 0 && y < m && x >= 0 && x < n && maze[y][x] == 0 {
				y, x, dist = y+dir[0], x+dir[1], dist+1
			}
			y, x, dist = y-dir[0], x-dir[1], dist-1
			if t[y][x] > dist {
				t[y][x] = dist
				if y != destination[0] || x != destination[1] {
					q = append(q, [2]int{y, x})
				}
			}
		}
	}
	out := t[destination[0]][destination[1]]
	if out == math.MaxInt64 {
		return -1
	}
	return out
}
