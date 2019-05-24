func longestPalindrome(s string) string {
	if 0 == len(s) {
		return ""
	}
	t := []byte{'#'}
	for _, c := range []byte(s) {
		t = append(t, c)
		t = append(t, '#')
	}

	m, mi := 0, 0
	for i := 0; i < len(t); i++ {
		k := i
		if len(t)-i-1 < k {
			k = len(t) - i - 1
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
	var i int
	if mi&1 == 1 {
		i = mi/2 + 1 - m/2
	} else {
		i = mi/2 - m/2
	}
	return s[i : i+m-1]
}
