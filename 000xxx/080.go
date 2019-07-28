func removeDuplicates(nums []int) int {
	l := len(nums)
	if 0 == l {
		return 0
	}
	i, j, c := 1, 0, 0
	for ; i < l; i++ {
		if nums[i] == nums[j] {
			c++
			if c >= 2 {
				continue
			}
		} else {
			c = 0
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}
