func missingElement(nums []int, k int) int {
	o, i := 0, 0
	for ; i < len(nums)-1; i++ {
		n, n1 := nums[i], nums[i+1]
		o = n + k
		if o < n1 {
			return o
		}
		k -= n1 - n - 1
	}
	return nums[i] + k
}
