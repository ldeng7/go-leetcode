func getConcatenation(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}
