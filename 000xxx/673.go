func findNumberOfLIS(nums []int) int {
	o, m, l := 0, 0, len(nums)
	ls, cs := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		ls[i], cs[i] = 1, 1
	}
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			if ls[i] == ls[j]+1 {
				cs[i] += cs[j]
			} else if ls[i] < ls[j]+1 {
				ls[i], cs[i] = ls[j]+1, cs[j]
			}
		}
		if m == ls[i] {
			o += cs[i]
		} else if m < ls[i] {
			m, o = ls[i], cs[i]
		}
	}
	return o
}
