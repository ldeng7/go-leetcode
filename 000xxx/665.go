func checkPossibility(nums []int) bool {
	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if c == 0 {
				return false
			} else if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
			c--
		}
	}
	return true
}
