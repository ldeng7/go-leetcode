func repeatedSubstringPattern(s string) bool {
	i, j, l := 1, 0, len(s)
	t := make([]int, l+1)
	for i < l {
		if s[i] == s[j] {
			i, j = i+1, j+1
			t[i] = j
		} else if j == 0 {
			i++
		} else {
			j = t[j]
		}
	}
	return t[l] != 0 && (t[l]%(l-t[l]) == 0)
}
