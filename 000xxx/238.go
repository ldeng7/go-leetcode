func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		out[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		out[i] = out[i-1] * nums[i-1]
	}
	t := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= t
		t *= nums[i]
	}
	return out
}
