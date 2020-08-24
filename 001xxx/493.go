func longestSubarray(nums []int) int {
	var o, c, p int
	b := false
	for _, n := range nums {
		if 0 != n {
			c++
			if t := p + c; t > o {
				o = t
			}
		} else {
			c, p, b = 0, c, true
		}
	}
	if !b {
		o--
	}
	return o
}
