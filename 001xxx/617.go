func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	o := make([]int, n-1)
	t := make([][]int, n)
	for _, e := range edges {
		a, b := e[0]-1, e[1]-1
		t[a], t[b] = append(t[a], b), append(t[b], a)
	}

	var cal func(int, int) [][2]int
	cal = func(u, p int) [][2]int {
		ars := make([][][2]int, 0, n)
		for _, v := range t[u] {
			if v != p {
				ars = append(ars, cal(v, u))
			}
		}
		ar := make([][2]int, 0, 16)

		var gen func(int, int, int, int)
		gen = func(i, m1, m2, d0 int) {
			if i == len(ars) {
				h, d := 0, 0
				if m1 < 0 {
				} else if m2 < 0 {
					h, d = m1+1, max(d0, m1+1)
				} else {
					h, d = m1+1, max(d0, m1+m2+2)
				}
				if d > 0 {
					o[d-1]++
				}
				ar = append(ar, [2]int{h, d})
			} else {
				for _, p := range ars[i] {
					h, d := p[0], p[1]
					mh1, mh2 := m1, m2
					if h > mh1 {
						mh2, mh1 = mh1, h
					} else if h > mh2 {
						mh2 = h
					}
					gen(i+1, mh1, mh2, max(d0, d))
				}
				gen(i+1, m1, m2, d0)
			}
		}
		gen(0, -1, -1, -1)
		return ar
	}
	cal(0, -1)
	return o
}
