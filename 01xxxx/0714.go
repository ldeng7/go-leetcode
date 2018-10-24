func maxProfit(prices []int, fee int) int {
	s, k := 0, -prices[0]
	for _, p := range prices {
		t := s
		if k+p-fee > s {
			s = k + p - fee
		}
		if t-p > k {
			k = t - p
		}
	}
	return s
}
