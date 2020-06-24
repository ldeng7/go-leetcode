var rules = [6][5]int{
	{1, 2, 3, 3, 4}, {3, 4, 1, 1, 2}, {1, 4, 2, 4, 1},
	{1, 3, 2, 3, 1}, {2, 4, 1, 4, 2}, {2, 3, 1, 3, 2},
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	if 1 == m && 1 == n {
		return true
	}

	var cal func(int, int, int) bool
	cal = func(y, x, pd int) bool {
		if y < 0 || y == m || x < 0 || x == n {
			return false
		}
		rule := rules[grid[y][x]-1]
		if pd == rule[0] || pd == rule[1] {
			return false
		}
		if y == m-1 && x == n-1 {
			return true
		}
		nd := rule[3]
		if pd != rule[2] {
			nd = rule[4]
		}
		switch nd {
		case 1:
			y++
		case 2:
			y--
		case 3:
			x++
		case 4:
			x--
		}
		return cal(y, x, nd)
	}

	g := grid[0][0]
	return ((g == 1 || g == 4 || g == 6) && cal(0, 1, 3)) || (g >= 2 && g <= 4 && cal(1, 0, 1))
}
