func maxProfit(prices []int) int {
	if 0 == len(prices) {
		return 0
	}
	b, out := prices[0], 0
	for _, p := range prices {
		if p < b {
			b = p
		}
		if p-b > out {
			out = p - b
		}
	}
	return out
}
