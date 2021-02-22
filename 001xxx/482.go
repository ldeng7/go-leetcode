func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	var l, r int = 1e9, 0
	for _, v := range bloomDay {
		l, r = min(v, l), max(v, r)
	}
	for l != r {
		c, n, d := 0, 0, l+(r-l)>>1
		for _, v := range bloomDay {
			if v <= d {
				c++
			} else {
				c = 0
			}
			if c == k {
				c, n = 0, n+1
			}
		}
		if n >= m {
			r = d
		} else {
			l = d + 1
		}
	}
	return l
}
