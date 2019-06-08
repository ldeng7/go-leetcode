func findDuplicate(nums []int) int {
	l, r := 1, len(nums)-1
	for l < r-1 {
		m := l + (r-l)>>1
		c := 0
		for _, num := range nums {
			if num < l || num > r {
				continue
			}
			if num <= m {
				c++
			} else {
				c--
			}
		}
		if c > 0 {
			r = m
		} else {
			l = m
		}
	}
	c := 0
	for _, num := range nums {
		if l == num {
			c++
		}
	}
	if c > 1 {
		return l
	}
	return r
}
