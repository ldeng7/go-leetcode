func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxAbsoluteSum(nums []int) int {
	s, mi, ma := 0, 0, 0
	for _, n := range nums {
		s += n
		mi, ma = min(mi, s), max(ma, s)
	}
	return ma - mi
}
