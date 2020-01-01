func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	l, ms := len(s), minSize-1
	var h, p uint64 = 0, 1
	for i := ms; i != 0; i-- {
		p *= 26
	}
	var o, b, d int
	c, m := [26]int{}, map[uint64]int{}

	for i := 0; i < l; i++ {
		j := s[i] - 'a'
		c[j]++
		if c[j] == 1 {
			d++
		}
		if i >= ms && i > ms+b {
			k := s[b] - 'a'
			c[k]--
			if c[k] == 0 {
				d--
			}
			h -= uint64(k) * p
			b++
		}
		h = h*26 + uint64(j)
		if i >= ms && d <= maxLetters {
			m[h]++
			if t := m[h]; t > o {
				o = t
			}
		}
	}
	return o
}
