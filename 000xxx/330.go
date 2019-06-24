func minPatches(nums []int, n int) int {
	o, c, i := 0, 1, 0
	for c <= n {
		if i < len(nums) && nums[i] <= c {
			c, i = c+nums[i], i+1
		} else {
			c, o = c<<1, o+1
		}
	}
	return o
}
