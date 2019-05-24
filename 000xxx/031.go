func nextPermutation(nums []int) {
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for k := 0; k < (len(nums)-i-1)>>1; k++ {
		nums[i+1+k], nums[len(nums)-1-k] = nums[len(nums)-1-k], nums[i+1+k]
	}
}
