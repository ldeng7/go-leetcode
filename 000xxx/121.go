func maxProfit(prices []int) int {
	if 0 == len(prices) {
		return 0
	}
	b, o := prices[0], 0
	for _, p := range prices {
		if p < b {
			b = p
		}
		if t := p - b; t > o {
			o = t
		}
	}
	return o
}
