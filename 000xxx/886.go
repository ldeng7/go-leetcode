func possibleBipartition(N int, dislikes [][]int) bool {
	g := make([][]int, N)
	for i := 0; i < N; i++ {
		g[i] = []int{}
	}
	m := map[int]bool{}
	for _, d := range dislikes {
		g[d[0]-1] = append(g[d[0]-1], d[1]-1)
		g[d[1]-1] = append(g[d[1]-1], d[0]-1)
	}

	var cal func(int, bool) bool
	cal = func(i int, b bool) bool {
		if b1, ok := m[i]; ok {
			return b1 == b
		}
		m[i] = b
		for _, j := range g[i] {
			if !cal(j, !b) {
				return false
			}
		}
		return true
	}

	for i := 0; i < N; i++ {
		if _, ok := m[i]; !ok && !cal(i, false) {
			return false
		}
	}
	return true
}
