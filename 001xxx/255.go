func maxScoreWords(words []string, letters []byte, score []int) int {
	cs := [26]int{}
	for _, l := range letters {
		cs[l-'a']++
	}
	o, l := 0, len(words)

	var cal func(int, int)
	cal = func(s, i int) {
		for ; i < l; i++ {
			j, sum, w := 0, 0, words[i]
			for ; j < len(w); j++ {
				k := w[j] - 'a'
				if 0 == cs[k] {
					break
				}
				cs[k]--
				sum += score[k]
			}
			if j == len(w) {
				cal(s+sum, i+1)
			}
			for k := 0; k < j; k++ {
				cs[w[k]-'a']++
			}
		}
		if s > o {
			o = s
		}
	}

	cal(0, 0)
	return o
}
