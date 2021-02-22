func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i2, j2 := -1, -1
	var s2 string
	for _, s1 := range word1 {
		l1 := len(s1)
		for j1 := 0; j1 < l1; j1++ {
			j2++
			if i2 == -1 || j2 == len(s2) {
				i2, j2 = i2+1, 0
				if i2 == len(word2) {
					return false
				}
				s2 = word2[i2]
			}
			if s1[j1] != s2[j2] {
				return false
			}
		}
	}
	return i2 == len(word2)-1 && j2 == len(s2)-1
}
