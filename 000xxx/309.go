import "math"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	b, pb, s, ps := math.MinInt64, 0, 0, 0
	for _, p := range prices {
		b, pb = max(b, ps-p), b
		s, ps = max(s, pb+p), s
	}
	return s
}
