func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	n <<= 1
	ar := make([]int, n)
	for i := 2; i < n; i++ {
		ar[i] = n + 1
	}
	t := make([][]int, n)
	for _, e := range redEdges {
		i := e[0] << 1
		t[i] = append(t[i], e[1]<<1+1)
	}
	for _, e := range blueEdges {
		i := e[0]<<1 + 1
		t[i] = append(t[i], e[1]<<1)
	}
	v := make([]int, n)
	v[0], v[1] = 1, 1
	q := []int{0, 1}

	for 0 != len(q) {
		i := q[0]
		q = q[1:]
		for _, j := range t[i] {
			ar[j] = min(ar[j], ar[i]+1)
			if 0 != v[j] {
				continue
			}
			v[j], q = 1, append(q, j)
		}
	}

	m := n >> 1
	o := make([]int, m)
	for i := 0; i < m; i++ {
		a := min(ar[i<<1], ar[i<<1+1])
		if a == n+1 {
			a = -1
		}
		o[i] = a
	}
	return o
}
