func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	j, o := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if k, ok := m[c]; ok && k >= j {
			j = k + 1
		}
		m[c] = i
		if o1 := i - j + 1; o1 > o {
			o = o1
		}
	}
	return o
}
