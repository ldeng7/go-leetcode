func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	o := 0
	i, j := 0, len(height)-1
	for i != j {
		hi, hj := height[i], height[j]
		s := (j - i) * min(hi, hj)
		if s > o {
			o = s
		}

		if hi > hj {
			j--
		} else {
			i++
		}
	}
	return o
}
