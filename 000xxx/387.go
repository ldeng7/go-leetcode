func firstUniqChar(s string) int {
	m := [26]int{}
	bs := []byte(s)
	for _, c := range bs {
		m[c-'a']++
	}
	for i, c := range bs {
		if 1 == m[c-'a'] {
			return i
		}
	}
	return -1
}
