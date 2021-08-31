func beautySum(s string) int {
	n, l := 1, len(s)
	m := [26]int{}
	bs := make([]int, l)
	for i := 0; i < l; i++ {
		c := s[i] - 'a'
		if m[c] == 0 {
			m[c] = n
			n++
		}
		bs[i] = m[c]
	}

	o := 0
	for i := 0; i < l; i++ {
		c, mi, ma := 1, 1, 1
		fs := make([]int, n)
		fs[bs[i]] = 1
		for j := i - 1; j >= 0; j-- {
			t := fs[bs[j]]
			fs[bs[j]]++
			if t == ma {
				ma++
			}
			if t == 0 {
				if mi == 1 {
					c++
				} else {
					c, mi = 1, 1
				}
			} else if t == mi {
				c--
				if c == 0 {
					c, mi = 0, mi+1
					for k := 0; k < n; k++ {
						if fs[k] == mi {
							c++
						}
					}
				}
			}
			o += ma - mi
		}
	}
	return o
}
