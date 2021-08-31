func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	o, l, l1 := 0, len(s), len(word1)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		t[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				t[i][j] = t[i-1][j+1] + 2
				if j < l1 && i >= l1 {
					o = max(o, t[i][j])
				}
			} else {
				t[i][j] = max(t[i-1][j], t[i][j+1])
			}
		}
	}
	return o
}
