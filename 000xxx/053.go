func maxSubArray(nums []int) int {
	var s *int
	c := 0
	for _, n := range nums {
		if c+n > n {
			c = c + n
		} else {
			c = n
		}
		if s != nil {
			if c > *s {
				*s = c
			}
		} else {
			c1 := c
			s = &c1
		}
	}
	return *s
}