func removeDuplicates(nums []int) int {
	l := len(nums)
	if 0 == l {
		return 0
	}
	i, j := 1, 0
	for ; i < l; i++ {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}
