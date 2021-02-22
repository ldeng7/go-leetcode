func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func minCost(s string, cost []int) int {
	o := 0
	for i, m := 0, 0; i < len(s); i++ {
		if i != 0 && s[i] != s[i-1] {
			m = 0
		}
		c := cost[i]
		o += min(m, c)
		m = max(m, c)
	}
	return o
}
