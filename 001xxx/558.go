func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func minOperations(nums []int) int {
	o, m := 0, 1
	for _, n := range nums {
		c := 0
		for ; n > 0; c, n = c+1, n>>1 {
			o += n & 1
		}
		m = max(m, c)
	}
	return o + m - 1
}
