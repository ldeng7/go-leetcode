func maxArea(height []int) int {
	out := 0
	i, j := 0, len(height)-1
	for i != j {
		h := height[i]
		if h > height[j] {
			h = height[j]
		}
		s := (j - i) * h
		if s > out {
			out = s
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return out
}
