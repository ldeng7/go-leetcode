func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	o := m
	t := make([][]bool, m)
	for i := 0; i < m; i++ {
		bs := make([]bool, n)
		for _, l := range languages[i] {
			bs[l-1] = true
		}
		t[i] = bs
	}

	bs := make([]bool, m+1)
	for _, f := range friendships {
		f0, f1, b := f[0]-1, f[1]-1, false
		bs1 := t[f1]
		for _, l := range languages[f0] {
			if bs1[l-1] {
				b = true
				break
			}
		}
		if !b {
			bs[f0], bs[f1] = true, true
		}
	}

	ar := make([]int, 0, m)
	for i, b := range bs {
		if b {
			ar = append(ar, i)
		}
	}
	for l := 0; l < n; l++ {
		c := 0
		for _, i := range ar {
			if !t[i][l] {
				c++
			}
		}
		o = min(o, c)
	}
	return o
}
