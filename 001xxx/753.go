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

func maximumScore(a int, b int, c int) int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, b, c = c, a, b
	} else if b > c {
		b, c = c, b
	}

	o := 0
	if t := a + b - c; t > 0 {
		t >>= 1
		o, a, b = o+t, a-t, b+t
	}
	return o + max(a, b) + min(a, c-b)
}
