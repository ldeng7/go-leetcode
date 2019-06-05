func hammingDistance(x int, y int) int {
	o, e := 0, x^y
	for e != 0 {
		o++
		e &= (e - 1)
	}
	return o
}
