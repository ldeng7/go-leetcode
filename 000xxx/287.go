func findDuplicate(nums []int) int {
	if 0 == len(nums) {
		return -1
	}
	l, r := 1, len(nums)-1
	for l < r {
		c, m := 0, l+(r-l)>>1
		for _, n := range nums {
			if n <= m {
				c++
			}
		}
		if c > m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
