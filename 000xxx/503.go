func nextGreaterElements(nums []int) []int {
	l := len(nums)
	out := make([]int, l)
	for i := 0; i < l; i++ {
		out[i] = -1
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < i+l; j++ {
			if nums[j%l] > nums[i] {
				out[i] = nums[j%l]
				break
			}
		}
	}
	return out
}
