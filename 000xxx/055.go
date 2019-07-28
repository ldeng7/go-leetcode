func canJump(nums []int) bool {
	l, j := len(nums), 0
	for i := 0; i < l; i++ {
		if i > j || j >= l-1 {
			break
		}
		if i+nums[i] > j {
			j = i + nums[i]
		}
	}
	return j >= l-1
}
