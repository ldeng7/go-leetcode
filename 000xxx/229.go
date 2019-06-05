func majorityElement(nums []int) []int {
	out := []int{}
	a, b, c1, c2, l := 0, 0, 0, 0, len(nums)
	for _, num := range nums {
		if num == a {
			c1++
		} else if num == b {
			c2++
		} else if c1 == 0 {
			a, c1 = num, 1
		} else if c2 == 0 {
			b, c2 = num, 1
		} else {
			c1, c2 = c1-1, c2-1
		}
	}
	c1, c2 = 0, 0
	for _, num := range nums {
		if num == a {
			c1++
		} else if num == b {
			c2++
		}
	}
	if c1 > l/3 {
		out = append(out, a)
	}
	if c2 > l/3 {
		out = append(out, b)
	}
	return out
}
