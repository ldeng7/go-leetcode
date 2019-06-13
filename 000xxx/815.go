func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	m := map[int]map[int]bool{}
	q := [][2]int{{S, 0}}
	for i, r := range routes {
		for _, j := range r {
			if sj, ok := m[j]; ok {
				sj[i] = true
			} else {
				m[j] = map[int]bool{i: true}
			}
		}
	}

	s := map[int]bool{}
	for 0 != len(q) {
		i, c := q[0][0], q[0][1]
		q = q[1:]
		if i == T {
			return c
		}
		for j, _ := range m[i] {
			for _, k := range routes[j] {
				if s[k] {
					continue
				}
				s[k] = true
				q = append(q, [2]int{k, c + 1})
			}
		}
	}
	return -1
}
