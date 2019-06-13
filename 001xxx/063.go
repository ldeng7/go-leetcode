func validSubarrays(nums []int) int {
	o, l := 0, len(nums)
	t := make([]int, l)
	for i := 0; i < l; i++ {
		t[i] = 1
	}
	for i := l - 1; i >= 0; i-- {
		j := i + 1
		for j < l && nums[j] >= nums[i] {
			j += t[j]
		}
		t[i] = j - i
		o += t[i]
	}
	return o
}
