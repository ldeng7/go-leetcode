func minStartValue(nums []int) int {
	s, m := 0, 0
	for _, n := range nums {
		s += n
		if s < m {
			m = s
		}
	}
	return -m + 1
}
