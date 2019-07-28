func longestWPI(hours []int) int {
	o, c, m := 0, 0, map[int]int{}
	for i, h := range hours {
		if h > 8 {
			c++
		} else {
			c--
		}
		if c > 0 {
			o = i + 1
		} else {
			if _, ok := m[c]; !ok {
				m[c] = i
			}
			if j, ok := m[c-1]; ok && i-j > o {
				o = i - j
			}
		}
	}
	return o
}
