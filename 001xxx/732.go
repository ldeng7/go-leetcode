func largestAltitude(gain []int) int {
	o, n := 0, 0
	for _, g := range gain {
		n += g
		if n > o {
			o = n
		}
	}
	return o
}
