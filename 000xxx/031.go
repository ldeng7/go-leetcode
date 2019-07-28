func nextPermutation(nums []int) {
	l := len(nums)
	i, j := l-2, l-1
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for k := 0; k < (l-i-1)>>1; k++ {
		nums[i+1+k], nums[l-1-k] = nums[l-1-k], nums[i+1+k]
	}
}
