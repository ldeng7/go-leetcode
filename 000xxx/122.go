func maxProfit(prices []int) int {
	l := len(prices)
	if 0 == l {
		return 0
	}
	prices = append(prices, 0)
	b, s, o := prices[0], 0, 0
	for i := 1; i <= l; i++ {
		p := prices[i]
		if p < prices[i-1] {
			o, b, s = o+s, p, 0
		}
		if t := p - b; t > s {
			s = t
		}
	}
	return o
}
