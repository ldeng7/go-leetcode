func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func numTimesAllBlue(light []int) int {
	o, m := 0, 0
	for i, l := range light {
		m = max(m, l)
		if m == i+1 {
			o++
		}
	}
	return o
}
