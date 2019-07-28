func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l>>1; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	if k <= 0 || k == l {
		return
	}
	if k > l {
		rotate(nums, k%l)
		return
	} else if l <= 1 {
		return
	}
	i := l - k
	reverse(nums[:i])
	reverse(nums[i:])
	reverse(nums)
}
