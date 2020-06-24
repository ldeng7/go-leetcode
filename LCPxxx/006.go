func minCount(coins []int) int {
	o := 0
	for _, n := range coins {
		o += (n + 1) >> 1
	}
	return o
}
