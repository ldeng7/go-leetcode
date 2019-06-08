func minEatingSpeed(piles []int, H int) int {
	var l, r int = 1, 1e9
	for l < r {
		m := l + (r-l)>>1
		c := 0
		for _, pile := range piles {
			c += (pile + m - 1) / m
		}
		if c > H {
			l = m + 1
		} else {
			r = m
		}
	}
	return r
}
