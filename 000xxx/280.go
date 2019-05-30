func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		if (i&1 == 1 && nums[i] < nums[i-1]) || (i&1 == 0 && nums[i] > nums[i-1]) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}
