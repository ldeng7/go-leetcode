func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		t := nums[len(nums)-1]
		for j := len(nums) - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = t
	}
}
