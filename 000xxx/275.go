func hIndex(citations []int) int {
	le := len(citations)
	l, r := 0, le-1
	for l <= r {
		m := l + (r-l)>>1
		if citations[m] == le-m {
			return le - m
		} else if citations[m] > le-m {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return le - l
}
