func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func getMaxLen(nums []int) int {
	o, c, f, j := 0, 0, -1, -1
	for i, n := range nums {
		if n < 0 {
			c++
			if f == -1 {
				f = i
			}
		}
		if n != 0 {
			if c&1 == 0 {
				o = max(i-j, o)
			} else {
				o = max(i-f, o)
			}
		} else {
			c, f, j = 0, -1, i
		}
	}
	return o
}
