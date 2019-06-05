func findMaxAverage(nums []int, k int) float64 {
	s, i := 0, 0
	for ; i < k; i++ {
		s += nums[i]
	}
	o := s
	for ; i < len(nums); i++ {
		s += nums[i] - nums[i-k]
		if s > o {
			o = s
		}
	}
	return float64(o) / float64(k)
}
