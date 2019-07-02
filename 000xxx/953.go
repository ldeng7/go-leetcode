func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func isAlienSorted(words []string, order string) bool {
	t := [26]int{}
	for i, c := range []byte(order) {
		t[c-'a'] = i
	}
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		c, l := 0, min(len(w1), len(w2))
		for j := 0; j < l; j++ {
			idx1, idx2 := t[w1[j]-'a'], t[w2[j]-'a']
			if idx1 > idx2 {
				return false
			} else if idx1 < idx2 {
				break
			}
			c++
		}
		if c == l && len(w1) > len(w2) {
			return false
		}
	}
	return true
}
