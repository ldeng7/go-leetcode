func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func criticalConnections(n int, connections [][]int) [][]int {
	g := make([][]int, n)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	o := [][]int{}
	v := make([]bool, n)
	low := make([]int, n)
	for i := 0; i < n; i++ {
		low[i] = -1
	}
	ids := make([]int, n)

	var cal func(int, int)
	cal = func(u, p int) {
		v[u] = true
		id := ids[p] + 1
		ids[u], low[u], v[u] = id, id, true
		for _, n := range g[u] {
			if n == p {
				continue
			}
			if !v[n] {
				cal(n, u)
				low[u] = min(low[u], low[n])
				if ids[u] < low[n] {
					o = append(o, []int{u, n})
				}
			} else {
				low[u] = min(low[u], ids[n])
			}
		}
	}

	for i := 0; i < n; i++ {
		if !v[i] {
			cal(i, i)
		}
	}
	return o
}
