func minTime(time []int, m int) int {
	if len(time) <= m {
		return 0
	}
	var l, r, mid int
	for _, t := range time {
		r += t
	}
	check := func() bool {
		var s, c, ma int
		for _, t := range time {
			s += t
			if t > ma {
				ma = t
			}
			if s-ma > mid {
				c++
				if c == m {
					return false
				}
				s, ma = t, t
			}
		}
		return true
	}

	for l < r {
		mid = l + (r-l)>>1
		if check() {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
