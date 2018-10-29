func findUnsortedSubarray(nums []int) int {
	l, s, e := len(nums), -1, -2
	min, max := nums[l-1], nums[0]
	for i := 1; i < l; i++ {
		j := l - 1 - i
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < max {
			e = i
		}
		if nums[j] < min {
			min = nums[j]
		} else if nums[j] > min {
			s = j
		}
	}
	return e - s + 1
}
