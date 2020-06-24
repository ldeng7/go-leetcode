func finalPrices(prices []int) []int {
	l := len(prices)
	o := make([]int, l)
	for i, p := range prices {
		d := 0
		for j := i + 1; j < l; j++ {
			if p1 := prices[j]; p1 <= p {
				d = p1
				break
			}
		}
		o[i] = p - d
	}
	return o
}
