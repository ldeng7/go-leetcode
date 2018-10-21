func moveZeroes(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
