func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}
	j, out := 0, 0
	for i, r := range s {
		if k, ok := m[r]; ok && k >= j {
			j = k + 1
		}
		m[r] = i

		l := i - j + 1
		if l > out {
			out = l
		}
	}
	return out
}
