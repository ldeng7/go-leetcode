func productExceptSelf(nums []int) []int {
	l := len(nums)
	out := make([]int, l)
	for i := 0; i < l; i++ {
		out[i] = 1
	}
	for i := 1; i < l; i++ {
		out[i] = out[i-1] * nums[i-1]
	}
	t := 1
	for i := l - 1; i >= 0; i-- {
		out[i] *= t
		t *= nums[i]
	}
	return out
}
