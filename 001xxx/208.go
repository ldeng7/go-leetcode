func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func equalSubstring(s string, t string, maxCost int) int {
	o, c, l := 0, 0, len(s)
	for i, j := 0, 0; i < l; i++ {
		c += abs(int(s[i]) - int(t[i]))
		for c > maxCost && j <= i {
			c -= abs(int(s[j]) - int(t[j]))
			j++
		}
		o = max(o, i-j+1)
	}
	return o
}
