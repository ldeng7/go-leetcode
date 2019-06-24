func canPartition(nums []int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	if s&1 == 1 {
		return false
	}
	t := make([]bool, s+1)
	s >>= 1
	t[0] = true
	for _, n := range nums {
		for i := s; i >= n; i-- {
			t[i] = t[i] || t[i-n]
		}
	}
	return t[s]
}
