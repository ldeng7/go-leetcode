func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func characterReplacement(s string, k int) int {
	o, c, b := 0, 0, 0
	cs := [26]int{}
	for i := 0; i < len(s); i++ {
		j := s[i] - 'A'
		cs[j]++
		c = max(c, cs[j])
		for i-b+1-c > k {
			cs[s[b]-'A']--
			b++
		}
		o = max(o, i-b+1)
	}
	return o
}
