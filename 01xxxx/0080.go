func removeDuplicates(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	i, j, c := 1, 0, 0
	for ; i < len(nums); i++ {
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
