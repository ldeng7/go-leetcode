func numIslands2(m int, n int, positions [][]int) []int {
	out := []int{}
	c := 0
	roots := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		roots[i] = -1
	}

	var getRoot func(int) int
	getRoot = func(i int) int {
		if i == roots[i] {
			return i
		}
		return getRoot(roots[i])
	}

	cal := func(py, px, dy, dx, i int) {
		y, x := py+dy, px+dx
		j := n*y + x
		if y < 0 || y >= m || x < 0 || x >= n || roots[j] == -1 {
			return
		}
		r, r1 := getRoot(i), getRoot(j)
		if r != r1 {
			roots[r1] = r
			c--
		}
	}

	for _, p := range positions {
		py, px := p[0], p[1]
		i := n*py + px
		if -1 == roots[i] {
			roots[i] = i
			c++
		}
		cal(py, px, -1, 0, i)
		cal(py, px, 1, 0, i)
		cal(py, px, 0, -1, i)
		cal(py, px, 0, 1, i)
		out = append(out, c)
	}
	return out
}
