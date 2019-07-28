func numSquares(n int) int {
	q := [][2]int{{0, 0}}
	v := make([]bool, n+1)
	v[0] = true
	for {
		for i := 1; ; i++ {
			tm, ti := q[0][0], q[0][1]
			m := tm + i*i
			if m > n {
				break
			}

			if m == n {
				return ti + 1
			}

			if !v[m] {
				q = append(q, [2]int{m, ti + 1})
				v[m] = true
			}
		}
		q = q[1:]
	}
}
