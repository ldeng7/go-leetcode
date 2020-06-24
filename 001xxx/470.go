func shuffle(nums []int, n int) []int {
	o := make([]int, len(nums))
	for i := 0; i < n; i++ {
		o[i*2], o[i*2+1] = nums[i], nums[i+n]
	}
	return o
}
