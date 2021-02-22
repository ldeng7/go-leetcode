func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maximumUniqueSubarray(nums []int) int {
	m := map[int]int{}
	o, s, j := 0, 0, 0
	for i, n := range nums {
		if v, ok := m[n]; ok {
			o = max(o, s)
			for ; j <= v; j++ {
				s -= nums[j]
			}
		}
		s += n
		m[n] = i
	}
	return max(o, s)
}
