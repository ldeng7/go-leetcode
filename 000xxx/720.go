func longestWord(words []string) string {
	ms := make([]map[string]bool, 30)
	for i := 0; i < 30; i++ {
		ms[i] = map[string]bool{}
	}
	for _, w := range words {
		ms[len(w)-1][w] = true
	}
	for i := 29; i >= 1; i-- {
		var wt string
		for w, _ := range ms[i] {
			ww := w
			j := i + 1
			for ; j >= 2; j-- {
				ww = ww[:j-1]
				if !ms[j-2][ww] {
					break
				}
			}
			if j == 1 && (len(wt) == 0 || w < wt) {
				wt = w
			}
		}
		if len(wt) > 0 {
			return wt
		}
	}
	var wt string
	for w, _ := range ms[0] {
		if len(wt) == 0 || w < wt {
			wt = w
		}
	}
	return wt
}
