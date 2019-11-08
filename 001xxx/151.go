func minSwaps(data []int) int {
	o, c, c1 := len(data), 0, 0
	for _, d := range data {
		if d == 1 {
			c++
		}
	}
	if c <= 1 {
		return 0
	}
	for i, d := range data {
		if d == 1 {
			c1++
		}
		if i >= c-1 {
			if t := c - c1; t < o {
				o = t
			}
			if data[i-c+1] == 1 {
				c1--
			}
		}
	}
	return o
}
