func canJump(nums []int) bool {
	j := 0
	for i := 0; i < len(nums); i++ {
		if i > j || j >= len(nums)-1 {
			break
		}
		if i+nums[i] > j {
			j = i + nums[i]
		}
	}
	return j >= len(nums)-1
}
