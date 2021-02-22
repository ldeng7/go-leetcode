type Edge struct {
	s, d, c int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maximumRequests(n int, requests [][]int) int {
	m, km := n+2, 0
	ds := make([]int, n)
	for _, r := range requests {
		ds[r[0]]++
		ds[r[1]]--
	}
	es := make([]Edge, 0, 16)
	for i, d := range ds {
		if d > 0 {
			for j := 0; j < d; j++ {
				es = append(es, Edge{n, i, 0})
			}
		} else if d < 0 {
			for j := 0; j < -d; j++ {
				es = append(es, Edge{i, n + 1, 0})
			}
		}
		km += max(d, 0)
	}
	for _, r := range requests {
		es = append(es, Edge{r[0], r[1], 1})
	}

	g := make([][]int, m)
	for i, e := range es {
		g[e.s], g[e.d] = append(g[e.s], i), append(g[e.d], i)
	}
	h := make([]int, m)
	o := len(requests)
	for k := 0; k < km; k++ {
		ds, ps := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			ds[i], ps[i] = m, -1
		}
		ds[n] = 0
		q, v := make([]int, 1, 16), make([]bool, m)
		q[0] = n

		for 0 != len(q) {
			t := q[0]
			q = q[1:]
			if v[t] {
				continue
			}
			v[t] = true
			for _, i := range g[t] {
				e := es[i]
				w := e.c + h[t] - h[e.d]
				if e.s != t || ds[e.d] <= ds[t]+w {
					continue
				}
				ds[e.d] = ds[t] + w
				if w != 0 {
					q = append(q, e.d)
				} else {
					q = append([]int{e.d}, q...)
				}
				ps[e.d] = i
			}
		}
		for i := 0; i < m; i++ {
			h[i] += ds[i]
		}
		o -= h[n+1]
		for i := n + 1; i != n; {
			e := &(es[ps[i]])
			e.s, e.d, e.c, i = e.d, e.s, -e.c, e.s
		}
	}
	return o
}
