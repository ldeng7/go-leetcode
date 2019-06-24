func checkSubarraySum(nums []int, k int) bool {
	s, m := 0, map[int]int{0: -1}
	for i, n := range nums {
		s += n
		t := s
		if k != 0 {
			t %= k
		}
		if j, ok := m[t]; ok {
			if i-j > 1 {
				return true
			}
		} else {
			m[t] = i
		}
	}
	return false
}
