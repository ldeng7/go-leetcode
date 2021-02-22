func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCharacters(a string, b string) int {
	m, n := len(a), len(b)
	ar1, ar2 := [26]int{}, [26]int{}
	for i := 0; i < m; i++ {
		ar1[a[i]-'a']++
	}
	for i := 0; i < n; i++ {
		ar2[b[i]-'a']++
	}
	o := m + n
	for i := 0; i < 26; i++ {
		o = min(o, m+n-ar1[i]-ar2[i])
	}
	for i, sa, sb := 0, 0, 0; i < 25; i++ {
		sa, sb = sa+ar1[i], sb+ar2[i]
		o = min(o, min(m-sa+sb, n+sa-sb))
	}
	return o
}
