func longestPalindrome(s string) string {
	l := len(s)
	if 0 == l {
		return ""
	}
	t, i := make([]byte, l<<1+1), 1
	t[0] = '#'
	for j := 0; j < l; j++ {
		t[i], t[i+1], i = s[j], '#', i+2
	}
	l = len(t)

	m, mi := 0, 0
	for i := 0; i < l; i++ {
		k := i
		if k1 := l - i - 1; k1 < k {
			k = k1
		}
		j := 1
		for ; j <= k; j++ {
			if t[i+j] != t[i-j] {
				break
			}
		}
		if j > m {
			m, mi = j, i
		}
	}
	if mi&1 == 1 {
		i = mi/2 + 1 - m/2
	} else {
		i = mi/2 - m/2
	}
	return s[i : i+m-1]
}
