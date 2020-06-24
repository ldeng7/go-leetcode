func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if b > a {
		a, b = b, a
	}
	for i := 2; i < len(nums); i++ {
		if c := nums[i]; c >= a {
			a, b = c, a
		} else if c > b {
			b = c
		}
	}
	return (a - 1) * (b - 1)
}
