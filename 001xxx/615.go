func maximalNetworkRank(n int, roads [][]int) int {
	ds := make([]int, n)
	for _, r := range roads {
		ds[r[0]]++
		ds[r[1]]++
	}
	m1, m2 := 0, 0
	for _, d := range ds {
		if d > m1 {
			m2, m1 = m1, d
		} else if d > m2 && d < m1 {
			m2 = d
		}
	}
	c1, c2 := 0, 0
	for _, d := range ds {
		if d == m1 {
			c1++
		}
		if d == m2 {
			c2++
		}
	}

	if c1 > 1 {
		c := 0
		for _, r := range roads {
			if ds[r[0]] == m1 && ds[r[1]] == m1 {
				c++
			}
		}
		a, b := 1, c1*(c1-1)/2
		if b != c {
			a = 0
		}
		return m1*2 - a
	}

	a, c := 0, 0
	for _, r := range roads {
		if ds[r[0]] == m1 && ds[r[1]] == m2 {
			c++
		}
		if ds[r[0]] == m2 && ds[r[1]] == m1 {
			c++
		}
	}
	if c == c2 {
		a = 1
	}
	return m1 + m2 - a
}
