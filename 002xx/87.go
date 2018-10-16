func findDuplicate(nums []int) int {
	l, h := 1, len(nums)-1
	for l < h-1 {
		m := int((uint(l) + uint(h)) >> 1)
		c := 0
		for _, num := range nums {
			if num < l || num > h {
				continue
			}
			if num <= m {
				c++
			} else {
				c--
			}
		}
		if c > 0 {
			h = m
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
	return h
}
