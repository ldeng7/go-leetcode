func longestSubstring(s string, k int) int {
	o, i, l := 0, 0, len(s)
	for i+k <= l {
		m, mi, cs := 0, i, [26]int{}
		for j := i; j < l; j++ {
			h := s[j] - 'a'
			cs[h]++
			if cs[h] < k {
				m |= (1 << h)
			} else {
				m &^= (1 << h)
			}
			if m == 0 {
				if o1 := j - i + 1; o1 > o {
					o = o1
				}
				mi = j
			}
		}
		i = mi + 1
	}
	return o
}
