func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestPalindromeSubseq(s string) int {
	o, l := 0, len(s)
	t := make([]int, l)
	for i := 0; i < l; i++ {
		t[i] = 1
	}
	for i := l - 1; i >= 0; i-- {
		n := 0
		for j := i + 1; j < l; j++ {
			t0 := t[j]
			if s[i] == s[j] {
				t[j] = n + 2
			}
			n = max(n, t0)
		}
	}
	for _, n := range t {
		o = max(o, n)
	}
	return o
}
