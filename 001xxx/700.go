func countStudents(students []int, sandwiches []int) int {
	c0, c1 := 0, 0
	for _, s := range students {
		if s == 0 {
			c0++
		} else {
			c1++
		}
	}
	for _, s := range sandwiches {
		if s == 0 {
			if c0 > 0 {
				c0--
			} else {
				break
			}
		} else {
			if c1 > 0 {
				c1--
			} else {
				break
			}
		}
	}
	return c0 + c1
}
