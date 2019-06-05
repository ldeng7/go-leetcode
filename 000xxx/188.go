func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if 0 == n {
		return 0
	}
	if k >= n {
		out := 0
		for i := 1; i < n; i++ {
			if prices[i]-prices[i-1] > 0 {
				out += prices[i] - prices[i-1]
			}
		}
		return out
	}
	g, l := make([]int, k+1), make([]int, k+1)
	for i := 0; i < n-1; i++ {
		d := prices[i+1] - prices[i]
		for j := k; j >= 1; j-- {
			l[j] = max(g[j-1]+max(d, 0), l[j]+d)
			g[j] = max(g[j], l[j])
		}
	}
	return g[k]
}
