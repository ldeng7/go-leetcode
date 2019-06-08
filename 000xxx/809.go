func expressiveWords(S string, words []string) int {
	o, m := 0, len(S)
	for _, w := range words {
		i, j, l := 0, 0, len(w)
		for ; i < m; i++ {
			if j < l && S[i] == w[j] {
				j++
			} else if i > 0 && S[i] == S[i-1] && i+1 < m && S[i] == S[i+1] {
				i++
			} else if i <= 1 || S[i] != S[i-1] || S[i] != S[i-2] {
				break
			}
		}
		if i == m && j == l {
			o++
		}
	}
	return o
}
