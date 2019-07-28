func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[endWord] {
		return [][]string{}
	}
	m := map[string][]string{}
	src, dst := map[string]bool{}, map[string]bool{}
	src[beginWord], dst[endWord] = true, true
	b, bw := false, false

	for 0 != len(src) && 0 != len(dst) && !b {
		if len(src) > len(dst) {
			src, dst, bw = dst, src, !bw
		}
		for w, _ := range src {
			dict[w] = false
		}
		src1 := map[string]bool{}
		for w, _ := range src {
			bs := []byte(w)
			for i := 0; i < len(bs); i++ {
				for j := byte('a'); j <= 'z'; j++ {
					bs[i] = j
					s, t := w, string(bs)
					if dst[t] {
						if bw {
							s, t = t, s
						}
						m[t], b = append(m[t], s), true
					} else if dict[t] {
						src1[t] = true
						if bw {
							s, t = t, s
						}
						m[t] = append(m[t], s)
					}
				}
				bs[i] = w[i]
			}
		}
		src = src1
	}

	o, ar := [][]string{}, []string{endWord}
	var cal func(string)
	cal = func(s string) {
		l := len(ar)
		if s == beginWord {
			ar1 := make([]string, l)
			copy(ar1, ar)
			for i := 0; i < l>>1; i++ {
				ar1[i], ar1[l-i-1] = ar1[l-i-1], ar1[i]
			}
			o = append(o, ar1)
			return
		}
		for _, t := range m[s] {
			ar = append(ar, t)
			cal(t)
			ar = ar[:l]
		}
	}
	cal(endWord)
	return o
}
