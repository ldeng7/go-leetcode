func threeEqualParts(A []int) []int {
	l, c1 := len(A), 0
	t := make([]int, l)
	for i, a := range A {
		if a == 1 {
			t[c1] = i
			c1++
		}
	}
	if c1 == 0 {
		return []int{0, l - 1}
	} else if c1%3 != 0 {
		return []int{-1, -1}
	}
	c := c1 / 3
	for i := 1; i < c; i++ {
		d := t[i] - t[i-1]
		if d != t[i+c]-t[i+c-1] || d != t[i+c<<1]-t[i+c<<1-1] {
			return []int{-1, -1}
		}
	}
	n := l - 1 - t[c1-1]
	if n > t[c<<1]-1-t[c<<1-1] || n > t[c]-1-t[c-1] {
		return []int{-1, -1}
	}
	return []int{t[c-1] + n, t[c<<1-1] + n + 1}
}
