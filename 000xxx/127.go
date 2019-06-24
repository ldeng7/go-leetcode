func ladderLength(beginWord string, endWord string, wordList []string) int {
	s := map[string]bool{}
	for _, w := range wordList {
		s[w] = true
	}
	if !s[endWord] {
		return 0
	}
	q := []string{beginWord}
	o := 0
	for 0 != len(q) {
		for i := len(q); i > 0; i-- {
			w := q[0]
			q = q[1:]
			if w == endWord {
				return o + 1
			}
			for j := 0; j < len(w); j++ {
				bs := []byte(w)
				for c := byte('a'); c <= 'z'; c++ {
					bs[j] = c
					str := string(bs)
					if s[str] && str != w {
						q = append(q, str)
						s[str] = false
					}
				}
			}
		}
		o++
	}
	return 0
}
