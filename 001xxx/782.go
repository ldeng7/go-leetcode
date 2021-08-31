func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func countPairs(n int, edges [][]int, queries []int) []int {
	ds, fs := make([]int, n), map[int]int{}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if u > v {
			u, v = v, u
		}
		ds[u]++
		ds[v]++
		fs[(u<<16)|v]++
	}

	md := ds[0]
	for i := 1; i < n; i++ {
		md = max(md, ds[i])
	}
	cs, dc := make([]int, md*2+2), map[int]int{}
	for i := 0; i < n; i++ {
		dc[ds[i]]++
	}
	for d1, c1 := range dc {
		for d2, c2 := range dc {
			if d1 < d2 {
				cs[d1+d2] += c1 * c2
			} else if d1 == d2 {
				cs[d1*2] += c1 * (c1 - 1) / 2
			}
		}
	}

	for k, f := range fs {
		u, v := k>>16, k&0xffff
		d := ds[u] + ds[v]
		cs[d]--
		cs[d-f]++
	}
	for i := len(cs) - 2; i >= 0; i-- {
		cs[i] += cs[i+1]
	}

	o := make([]int, len(queries))
	for i, q := range queries {
		o[i] = cs[min(q+1, len(cs)-1)]
	}
	return o
}
