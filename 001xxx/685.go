func getSumAbsoluteDifferences(nums []int) []int {
	l := len(nums)
	o := make([]int, l)
	s, n0 := 0, nums[0]
	for _, n := range nums {
		s += n - n0
	}
	o[0] = s
	for i := 1; i < l; i++ {
		o[i] += o[i-1] + (nums[i]-nums[i-1])*i - (nums[i]-nums[i-1])*(l-i)
	}
	return o
}
