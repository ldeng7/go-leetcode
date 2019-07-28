func canDivideIntoSubsequences(nums []int, K int) bool {
	c, m, np, l := 1, 0, nums[0], len(nums)
	for i := 1; i < l; i++ {
		n := nums[i]
		if n == np {
			c++
			if c > m {
				m = c
			}
		} else {
			c, np = 1, n
		}
	}
	return m*K <= l
}
