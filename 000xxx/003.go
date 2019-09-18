func lengthOfLongestSubstring(s string) int {
	o, b, e, i, j := 0, 0, -1, 0, 0
	for ; i < len(s); i++ {
		for j = b; j <= e; j++ {
			if s[j] == s[i] {
				break
			}
		}
		if j > e {
			e = i
			if t := e - b + 1; t > o {
				o = t
			}
		} else {
			b, e = j+1, i
		}
	}
	return o
}
