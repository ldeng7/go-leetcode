func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func numWays(words []string, target string) int {
	l, n := len(target), len(words[0])
	cs := make([][26]int, n)
	for _, w := range words {
		for j := 0; j < n; j++ {
			cs[j][w[j]-'a']++
		}
	}
	t := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		t[i] = make([]int, l+1)
		t[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		je := min(l, i)
		for j := 1; j <= je; j++ {
			t[i][j] = (t[i-1][j] + t[i-1][j-1]*cs[i-1][target[j-1]-'a']) % 1000000007
		}
	}
	return t[n][l]
}
