func moveZeroes(nums []int) {
	l, i, j := len(nums), 0, 0
	for ; i < l; i++ {
		if nums[i] == 0 {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	for ; j < l; j++ {
		nums[j] = 0
	}
}
