func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	lp := len(prices)
	if 0 == lp {
		return 0
	}
	g, l := [3]int{}, [3]int{}
	for i := 0; i < lp-1; i++ {
		d := prices[i+1] - prices[i]
		for j := 2; j >= 1; j-- {
			l[j] = max(g[j-1]+max(d, 0), l[j]+d)
			g[j] = max(l[j], g[j])
		}
	}
	return g[2]
}
