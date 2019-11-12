func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func isValidPalindrome(s string, k int) bool {
	l := len(s)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		ar := make([]int, l)
		for j := 0; j < l; j++ {
			ar[j] = 1
		}
		t[i] = ar
	}
	for i := 1; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if s[j] == s[j+i] {
				if i == 1 {
					t[j][j+i] = 2
				} else {
					t[j][j+i] = t[j+1][j+i-1] + 2
				}
			} else {
				t[j][j+i] = max(t[j][j+i-1], t[j+1][j+i])
			}
		}
	}
	return t[0][l-1]+k >= l
}
