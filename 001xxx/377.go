func frogPosition(n int, edges [][]int, t int, target int) float64 {
	n++
	g := make([][]int, n)
	v := make([]bool, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	var cal func(int, int) float64
	cal = func(i, t int) float64 {
		v[i] = true
		c := len(g[i])
		if i != 1 {
			c--
		}
		if i == target {
			if t == 0 || c == 0 {
				return 1
			}
			return 0
		} else if t == 0 {
			return 0
		}

		var o float64
		for _, j := range g[i] {
			if v[j] {
				continue
			}
			o = cal(j, t-1) / float64(c)
			if o != 0 {
				break
			}
		}
		return o
	}
	return cal(1, t)
}
