func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	var out, a int
	st := []int{}
	for i, h := range heights {
		j := len(st) - 1
		for -1 != j && h <= heights[st[j]] {
			t := st[j]
			st = st[:j]
			j--
			if -1 != j {
				a = heights[t] * (i - st[j] - 1)
			} else {
				a = heights[t] * i
			}
			if a > out {
				out = a
			}
		}
		st = append(st, i)
	}
	return out
}
