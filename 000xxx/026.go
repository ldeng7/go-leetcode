func removeDuplicates(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}
