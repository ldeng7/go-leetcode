func nthUglyNumber(n int) int {
	sl := []int{1}
	var c2, c3, c5 int
	for len(sl) < n {
		i2, i3, i5 := sl[c2]*2, sl[c3]*3, sl[c5]*5
		i := i2
		if i3 < i {
			i = i3
		}
		if i5 < i {
			i = i5
		}
		if i == i2 {
			c2++
		}
		if i == i3 {
			c3++
		}
		if i == i5 {
			c5++
		}
		sl = append(sl, i)
	}
	return sl[len(sl)-1]
}
