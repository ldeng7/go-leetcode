func lengthOfLongestSubstringKDistinct(s string, k int) int {
	o, l := 0, 0
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		for len(m) > k {
			cl := s[l]
			m[cl]--
			if m[cl] == 0 {
				delete(m, cl)
			}
			l++
		}
		o1 := i - l + 1
		if o1 > o {
			o = o1
		}
	}
	return o
}
