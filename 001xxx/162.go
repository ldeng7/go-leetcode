var dirs = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func maxDistance(grid [][]int) int {
	o, l := -1, len(grid)
	q := [][2]int{}
	for i, line := range grid {
		for j, g := range line {
			if 0 != g {
				q = append(q, [2]int{i, j})
			}
		}
	}
	if len(q) == l*l {
		return -1
	}
	for 0 != len(q) {
		n := len(q)
		for 0 != n {
			n--
			y, x := q[0][0], q[0][1]
			q = q[1:]
			for _, d := range dirs {
				r, c := y+d[0], x+d[1]
				if r < 0 || r >= l || c < 0 || c >= l || 0 != grid[r][c] {
					continue
				}
				q = append(q, [2]int{r, c})
				grid[r][c] = 1
			}
		}
		o++
	}
	return o
}
