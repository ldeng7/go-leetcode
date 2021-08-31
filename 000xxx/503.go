func nextGreaterElements(nums []int) []int {
	l := len(nums)
	o := make([]int, l)
	for i := 0; i < l; i++ {
		o[i] = -1
	}
	s := make([]int, 0, 16)
	for i := 0; i < l<<1; i++ {
		n := nums[i%l]
		for len(s) > 0 {
			j := len(s) - 1
			m := s[j]
			if n <= nums[m] {
				break
			}
			o[m] = n
			s = s[:j]
		}
		if i < l {
			s = append(s, i)
		}
	}
	return o
}
