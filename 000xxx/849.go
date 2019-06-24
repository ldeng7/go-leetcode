func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxDistToClosest(seats []int) int {
	b, o := 0, 0
	for i, s := range seats {
		if s != 1 {
			continue
		}
		if b == 0 {
			o = max(o, i-b)
		} else {
			o = max(o, (i-b+1)>>1)
		}
		b = i + 1
	}
	o = max(o, len(seats)-b)
	return o
}
