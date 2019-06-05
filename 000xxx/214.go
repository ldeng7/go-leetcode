func shortestPalindrome(s string) string {
	ls := len(s)
	lt := ls*2 + 1
	t := make([]byte, lt)
	copy(t, []byte(s))
	t[ls] = '#'
	i := ls + 1
	for j := ls - 1; j >= 0; j-- {
		t[i] = s[j]
		i++
	}

	next := make([]int, lt)
	for i := 1; i < lt; i++ {
		j := next[i-1]
		for j > 0 && t[i] != t[j] {
			j = next[j-1]
		}
		if t[i] == t[j] {
			j++
		}
		next[i] = j
	}

	lh := ls - next[lt-1]
	out := make([]byte, lh+ls)
	copy(out, t[ls+1:ls+1+lh])
	copy(out[lh:], t[:ls])
	return string(out)
}
