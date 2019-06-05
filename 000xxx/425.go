func wordSquares(words []string) [][]string {
	out := [][]string{}
	m := map[string]map[string]bool{}
	n := len(words[0])
	for _, w := range words {
		for i := 0; i < n; i++ {
			k := w[:i]
			if mk, ok := m[k]; ok {
				mk[w] = true
			} else {
				m[k] = map[string]bool{w: true}
			}
		}
	}
	mat := make([][]byte, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]byte, n)
	}

	var cal func(int, int)
	cal = func(i, n int) {
		if i == n {
			ar := []string{}
			for j := 0; j < n; j++ {
				ar = append(ar, string(mat[j]))
			}
			out = append(out, ar)
			return
		}
		k := string(mat[i][:i])
		for s, _ := range m[k] {
			mat[i][i] = s[i]
			j := i + 1
			for ; j < n; j++ {
				mat[i][j], mat[j][i] = s[j], s[j]
				if _, ok := m[string(mat[j][:i+1])]; !ok {
					break
				}
			}
			if j == n {
				cal(i+1, n)
			}
		}
	}
	cal(0, n)
	return out
}
