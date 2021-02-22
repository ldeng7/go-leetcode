func getIndex(reader *ArrayReader) int {
	var cal func(int, int, int, int) int
	cal = func(l, r, x, y int) int {
		if v := reader.compareSub(l, r, x, y); v > 0 {
			if l == r {
				return l
			}
			n := r - l + 1
			return cal(l, l+n/2-1, r-n/2+1, r)
		} else if v < 0 {
			if x == y {
				return x
			}
			n := y - x + 1
			return cal(x, x+n/2-1, y-n/2+1, y)
		}
		return r + 1
	}
	n := reader.length()
	return cal(0, n/2-1, n-n/2, n-1)
}
