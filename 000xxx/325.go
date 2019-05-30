func maxSubArrayLen(nums []int, k int) int {
	s, o := 0, 0
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		if s == k {
			o = i + 1
		} else if j, ok := m[s-k]; ok {
			if i-j > o {
				o = i - j
			}
		}
		if _, ok := m[s]; !ok {
			m[s] = i
		}
	}
	return o
}
