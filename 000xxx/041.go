func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
