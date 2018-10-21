func maxProfit(prices []int) int {
	out, t := 0, -1
	for _, p := range prices {
		if -1 == t || p < t {
			t = p
		}
		if p-t > out {
			out = p - t
		}
	}
	return out
}
