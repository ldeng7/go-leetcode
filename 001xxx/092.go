func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
		t[i][0] = i
	}
	for j := 1; j <= n; j++ {
		t[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				t[i][j] = t[i-1][j-1] + 1
			} else {
				t[i][j] = min(t[i-1][j], t[i][j-1]) + 1
			}
		}
	}

	i, j, bs := m, n, make([]byte, 0, m+n)
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			bs, i, j = append(bs, str1[i-1]), i-1, j-1
		} else if t[i-1][j] < t[i][j-1] {
			bs, i = append(bs, str1[i-1]), i-1
		} else {
			bs, j = append(bs, str2[j-1]), j-1
		}
	}
	l := len(bs)
	for k := 0; k < l>>1; k++ {
		bs[k], bs[l-k-1] = bs[l-k-1], bs[k]
	}
	return str2[:j] + str1[:i] + string(bs)
}
