func maxProfit(prices []int) int {
	if 0 == len(prices) {
		return 0
	}
	prices = append(prices, 0)
	b, s, out := prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		p := prices[i]
		if p < prices[i-1] {
			b = p
			out += s
			s = 0
		}
		if p-b > s {
			s = p - b
		}
	}
	return out
}
