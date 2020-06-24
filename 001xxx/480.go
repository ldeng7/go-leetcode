func runningSum(nums []int) []int {
	l := len(nums)
	o := make([]int, l)
	s := nums[0]
	o[0] = s
	for i := 1; i < l; i++ {
		s += nums[i]
		o[i] = s
	}
	return o
}
