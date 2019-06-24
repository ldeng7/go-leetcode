func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	m1, m2, m3, m4 := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	m := map[int]map[int]bool{}
	for _, lamp := range lamps {
		y, x := lamp[0], lamp[1]
		m1[y]++
		m2[x]++
		m3[y-x]++
		m4[y+x]++
		if my, ok := m[y]; ok {
			my[x] = true
		} else {
			m[y] = map[int]bool{x: true}
		}
	}
	o := make([]int, len(queries))
	for i, q := range queries {
		y, x := q[0], q[1]
		if m1[y] > 0 || m2[x] > 0 || m3[y-x] > 0 || m4[y+x] > 0 {
			o[i] = 1
		}
		a := y - 1
		if a < 0 {
			a++
		}
		for ; a <= y+1 && a < N; a++ {
			b := x - 1
			if b < 0 {
				b++
			}
			for ; b <= x+1 && b < N; b++ {
				if _, ok := m[a]; ok && m[a][b] {
					m1[a]--
					m2[b]--
					m3[a-b]--
					m4[a+b]--
					m[a][b] = false
				}
			}
		}
	}
	return o
}
