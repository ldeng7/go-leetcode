func minimumBoxes(n int) int {
	o, s, i := 0, 0, 1
	for ; i <= n && s < n; i++ {
		o = i * (i + 1) / 2
		s += o
	}
	if s >= n {
		i--
	}
	for ; i <= s-n; i-- {
		o, s = o-1, s-i
	}
	return o
}
